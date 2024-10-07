// Package main is the entrypoint of the application.
package main

import (
	"log"
	_ "template/docs"
	"template/internal/adapters/entrypoints/swagger"
	"template/internal/adapters/entrypoints/v1"
	"template/internal/adapters/mongodb"
	"template/internal/adapters/postgresql"
	containers "template/internal/dependency-injection"

	"github.com/gin-gonic/gin"
)

func main() {
	shutdown := postgresql.InitTracer()
	defer shutdown()

	// Create a MongoDB client
	clientMongo, err := mongodb.NewMongoClient("mongodb://root:example@localhost:27017")
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	clientPostgresql, err := postgresql.NewPostgresClient("localhost", "root", "example", "postgres_db", 5432)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	if err = postgresql.Migrate(clientPostgresql); err != nil {
		log.Fatalf("Could not migrations: %v", err)
	}

	// Instantiate the container with the MongoDB client
	domainContainer := containers.NewDomainContainer(clientMongo, clientPostgresql, "mongodb_data")
	entrypointContainer := containers.NewEntrypointContainer(domainContainer)

	// Create the Gin router
	engine := gin.Default()

	swagger.SetupSwagger(engine)

	entrypoints.RegisterRoutes(engine)
	// Register the routes using the container's controller
	entrypointContainer.V1EntityController.RegisterRoutes(engine)
	entrypointContainer.V1UserController.RegisterRoutes(engine)

	// Start the server
	if err := engine.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

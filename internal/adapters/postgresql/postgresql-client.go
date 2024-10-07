// Package postgresql contains the implementation to use postgres as the application database
package postgresql

import (
	"context"
	"fmt"
	"log"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/opentelemetry/tracing"
)

func InitTracer() func() {
	// Create stdout exporter to see the spans in the console
	exp, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Fatalf("failed to initialize stdouttrace exporter: %v", err)
	}

	// Create a new tracer provider with the exporter
	tp := sdktrace.NewTracerProvider(sdktrace.WithBatcher(exp))

	// Set the global tracer provider
	otel.SetTracerProvider(tp)

	// Return a function to shutdown the tracer provider
	return func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shut down tracer provider: %v", err)
		}
	}
}

// NewPostgresClient creates a new PostgreSQL client using GORM
func NewPostgresClient(host, user, password, dbname string, port int) (*gorm.DB, error) {
	// Formating connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", host, user, password, dbname, port)

	// Opening connection to PostgreSQL using GORM
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.Use(tracing.NewPlugin()); err != nil {
		return nil, err
	}

	// Optionally, you can ping the database to verify the connection
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

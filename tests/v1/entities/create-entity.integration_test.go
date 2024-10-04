package entities

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	v1 "template/internal/adapters/entrypoints/v1"
	"template/internal/adapters/mongodb"
	"template/internal/domain/entities"
	"template/internal/domain/usecases"

	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateEntityEndpoint(t *testing.T) {
	defer CleanupDatabase(t)
	// Given that the database is configured and active
	// And that all dependencies have been started
	repository := mongodb.NewEntityRepository(MongoClient, "test_db")
	useCase := usecases.NewCreateEntityUseCase(repository)
	controller := v1.NewEntityController(useCase, nil, nil, nil)
	// And that the correct parameters are passed to the entrypoint
	entity := entities.Entity{Name: "Test Entity"}
	body, _ := json.Marshal(entity)
	// And that the request is made to the correct URI
	req, err := http.NewRequest("POST", "/entities", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := gin.Default()
	controller.RegisterRoutes(router)

	// When calling the entity creation service
	router.ServeHTTP(rr, req)

	// Then the status code should be 201
	assert.Equal(t, http.StatusCreated, rr.Code)
	// And there should be no errors
	var createdEntity entities.Entity
	err = json.Unmarshal(rr.Body.Bytes(), &createdEntity)
	// And the response body should be empty
	assert.NoError(t, err)
	assert.Empty(t, createdEntity.ID)
	assert.Empty(t, createdEntity.Name)
}

func TestCreateEntityMissingName(t *testing.T) {
	defer CleanupDatabase(t)
	// Given that the database is configured and active
	// And that all dependencies have been started
	repository := mongodb.NewEntityRepository(MongoClient, "test_db")
	useCase := usecases.NewCreateEntityUseCase(repository)
	controller := v1.NewEntityController(useCase, nil, nil, nil)
	// And that the request body not contain the required name field
	body, _ := json.Marshal(map[string]string{"outher": "field"})
	// And that the request is made to the correct URI
	req, err := http.NewRequest("POST", "/entities", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	router := gin.Default()
	controller.RegisterRoutes(router)

	// When calling the entity creation service
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	// Faz o parse do corpo da resposta
	var responseBody map[string]interface{}
	responseError := json.Unmarshal(rr.Body.Bytes(), &responseBody)
	if responseError != nil {
		t.Fatalf("Erro ao fazer o parse do corpo da resposta: %v", responseError)
	}

	// Verifica se a mensagem de erro está correta
	assert.Contains(t, responseBody["error"], "EntityDTO.Name") // Exemplo de mensagem esperada
}

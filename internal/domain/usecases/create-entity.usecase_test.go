package usecases

import (
	"errors"
	"template/internal/domain/entities"
	"template/internal/domain/gateways/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEntityUseCaseExecute(t *testing.T) {
	// Given that I have a valid entity to be created
	mockName := "Test Entity"
	expectedEntity := entities.Entity{Name: mockName}
	// And the gateway will return a created entity with success
	mockGateway := new(mocks.EntityGateway)
	mockGateway.On("Create", expectedEntity).Return(entities.Entity{ID: "mocked-id", Name: mockName}, nil)

	// When calling the use case
	useCase := NewCreateEntityUseCase(mockGateway)
	createdEntity, err := useCase.Execute(expectedEntity)

	// Then the gateway method should have been called with the correct parameter
	mockGateway.AssertCalled(t, "Create", expectedEntity)
	// And the entity should be created successfully
	assert.NoError(t, err)
	assert.Equal(t, "mocked-id", createdEntity.ID)
	assert.Equal(t, mockName, createdEntity.Name)
}

func TestCreateEntityUseCaseExecuteWithErrorInGateway(t *testing.T) {
	// Given that I have a valid entity to be created
	mockName := "Test Entity"
	expectedEntity := entities.Entity{Name: mockName}
	// And the gateway will return an error
	mockGateway := new(mocks.EntityGateway)
	mockGateway.On("Create", expectedEntity).Return(entities.Entity{}, errors.New("creation error"))

	// When calling the use case
	useCase := NewCreateEntityUseCase(mockGateway)
	createdEntity, err := useCase.Execute(expectedEntity)

	// Then the gateway method should have been called with the correct parameter
	mockGateway.AssertCalled(t, "Create", expectedEntity)
	// And the entity should not be created
	assert.Error(t, err)
	assert.EqualError(t, err, "creation error")
	assert.Empty(t, createdEntity.ID)
	assert.Empty(t, createdEntity.Name)
}

// Package entrypoints contains the controllers that handle HTTP requests and call the use cases
package entrypoints

import (
	"net/http"
	"template/internal/adapters/entrypoints/v1/presenters"
	"template/internal/domain/usecases"

	"github.com/gin-gonic/gin"
)

// UserController handles the operations for the user
type UserController struct {
	createUseCase *usecases.CreateUserUseCase
	getUseCase    *usecases.GetUserUseCase
}

// NewUserController creates a new UserController
func NewUserController(
	createUseCase *usecases.CreateUserUseCase,
	getUseCase *usecases.GetUserUseCase,
) *UserController {
	return &UserController{
		createUseCase: createUseCase,
		getUseCase:    getUseCase,
	}
}

// RegisterRoutes registers the routes for the user
func (controller *UserController) RegisterRoutes(route *gin.Engine) {
	route.POST("/users", controller.CreateUser)
	route.GET("/users/:id", controller.GetUser)
}

// CreateUser handles the creation of a new user
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body presenters.UserDTO true "User details"
// @Success 201 {string} id the new user
// @Failure 400
// @Failure 500
// @Router /users [post]
func (controller *UserController) CreateUser(context *gin.Context) {
	var dto presenters.UserDTO
	if err := context.ShouldBindJSON(&dto); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity := presenters.MapperDTOToUser(dto)
	newUser, err := controller.createUseCase.Execute(entity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, newUser.ID)
}

// GetUser handles the retrieval of an entity by ID
// @Summary Get an entity by ID
// @Description Retrieve an entity by its ID
// @Tags users
// @Produce json
// @Param id path string true "Entity ID"
// @Success 200 {object} presenters.UserDTO
// @Failure 404
// @Router /users/{id} [get]
func (controller *UserController) GetUser(context *gin.Context) {
	id := context.Param("id")
	entity, err := controller.getUseCase.Execute(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	context.JSON(http.StatusOK, presenters.MapperUserToDTO(entity))
}

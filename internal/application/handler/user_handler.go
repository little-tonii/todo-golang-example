package handler

import (
	"todo-golang-example/internal/domain/repository"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepository repository.UserRepository
}

func NewUserHandler(userRepository repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepository: userRepository,
	}
}

func (handler *UserHandler) Register(context *gin.Context) {}

func (handler *UserHandler) Login(context *gin.Context) {}

func (handler *UserHandler) Info(context *gin.Context) {}

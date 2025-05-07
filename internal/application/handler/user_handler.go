package handler

import (
	"todo-golang-example/internal/domain/repository"
	repositoryimpl "todo-golang-example/internal/infrastructure/repository_impl"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userRepository repository.UserRepository
}

func NewUserHandler() *UserHandler {
	return &UserHandler{
		userRepository: repositoryimpl.GetUserRepository(),
	}
}

func (handler *UserHandler) Register(context *gin.Context) {}

func (handler *UserHandler) Login(context *gin.Context) {}

func (handler *UserHandler) Info(context *gin.Context) {}

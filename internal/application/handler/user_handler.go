package handler

import (
	"errors"
	"net/http"
	"todo-golang-example/internal/application/request"
	"todo-golang-example/internal/application/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (handler *UserHandler) HandleRegister() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestRaw, exists := context.Get("request_data")
		if !exists {
			context.Error(errors.New("Dữ liệu request bị mất sau khi validate"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, parseResult := requestRaw.(*request.RegisterUserRequest)
		if !parseResult {
			context.Error(errors.New("Dữ liệu request không phải là RegisterUserRequest"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		error := handler.userService.Register(request)
		if error != nil {
			context.Error(errors.New(error.Error()))
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Đăng ký thành công"})
	}
}

func (handler *UserHandler) HandleLogin() gin.HandlerFunc {
	return func(context *gin.Context) {}
}

func (handler *UserHandler) HandleInfo() gin.HandlerFunc {
	return func(context *gin.Context) {}
}

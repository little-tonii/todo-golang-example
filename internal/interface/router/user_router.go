package router

import (
	"todo-golang-example/internal/application/handler"
	"todo-golang-example/internal/application/request"
	"todo-golang-example/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeUserRouter(engine *gin.Engine) {
	userHandler := handler.NewUserHandler()

	userGroup := engine.Group("/user")
	{
		userGroup.POST("/register", middleware.BindingValidator[request.RegisterUserRequest](), userHandler.Register)
		userGroup.POST("/login", middleware.BindingValidator[request.LoginUserRequest](), userHandler.Login)
	}

	requiredAuthentication := userGroup.Group("")
	{
		requiredAuthentication.GET("/info", middleware.Authentication(), userHandler.Info)
	}
}

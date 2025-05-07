package router

import (
	"todo-golang-example/internal/application/handler"
	"todo-golang-example/internal/application/request"
	repositoryimpl "todo-golang-example/internal/infrastructure/repository_impl"
	"todo-golang-example/internal/shared/config"
	"todo-golang-example/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeUserRouter(engine *gin.Engine) {
	database := config.GetDatabase()
	userRepository := repositoryimpl.NewUserRepositoryImpl(database)
	userHandler := handler.NewUserHandler(userRepository)

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

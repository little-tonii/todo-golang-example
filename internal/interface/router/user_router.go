package router

import (
	"todo-golang-example/internal/application/request"
	"todo-golang-example/internal/application/service"
	"todo-golang-example/internal/infrastructure/config"
	repositoryimpl "todo-golang-example/internal/infrastructure/repository_impl"
	"todo-golang-example/internal/interface/handler"
	"todo-golang-example/pkg/middleware"

	"github.com/gin-gonic/gin"
)

func InitializeUserRouter(engine *gin.Engine) {
	database := config.GetDatabase()
	userRepository := repositoryimpl.NewUserRepositoryImpl(database)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	userGroup := engine.Group("/user")

	{
		userGroup.POST(
			"/register",
			middleware.BindingValidator[request.RegisterUserRequest](),
			userHandler.HandleRegister(),
		)
		userGroup.POST(
			"/login",
			middleware.BindingValidator[request.LoginUserRequest](),
			userHandler.HandleLogin(),
		)
	}

	requiredAuthentication := userGroup.Group("")

	{
		requiredAuthentication.GET(
			"/info",
			middleware.Authentication(),
			userHandler.HandleInfo(),
		)
	}
}

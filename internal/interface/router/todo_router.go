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

func InitializeTodoRouter(engine *gin.Engine) {
	database := config.GetDatabase()
	todoRepository := repositoryimpl.NewTodoRepositoryImpl(database)
	todoService := service.NewTodoService(todoRepository)
	todoHandler := handler.NewTodoHandler(todoService)

	todoGroup := engine.Group("/todo", middleware.Authentication())

	{
		todoGroup.GET("/:id", todoHandler.HandleGetTodoById())
		todoGroup.DELETE("/:id", todoHandler.HandlerDeleteTodoById())
		todoGroup.POST(
			"/create",
			middleware.BindingValidator[request.CreateTodoRequest](),
			todoHandler.HandleCreateTodo(),
		)
		todoGroup.GET(
			"/list",
			middleware.BindingValidator[request.GetTodoListRequest](),
			todoHandler.HandleGetTodoList(),
		)
		todoGroup.PUT(
			"/update/:id",
			middleware.BindingValidator[request.UpdateTodoByIdRequest](),
			todoHandler.HandleUpdateTodoById(),
		)
	}
}

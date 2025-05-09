package handler

import (
	"errors"
	"net/http"
	"strconv"
	"todo-golang-example/internal/application/request"
	"todo-golang-example/internal/application/service"
	"todo-golang-example/pkg/utils"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	todoService *service.TodoService
}

func NewTodoHandler(todoService *service.TodoService) *TodoHandler {
	return &TodoHandler{
		todoService: todoService,
	}
}

func (handler *TodoHandler) HandleGetTodoById() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, error := strconv.ParseInt(context.Param("id"), 10, 64)
		if error != nil || id <= 0 {
			context.Error(errors.New("Id phải là một số nguyên dương"))
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		response, appError := handler.todoService.GetTodoById(id)
		if appError != nil {
			context.Error(appError.Error)
			context.AbortWithStatus(appError.StatusCode)
			return
		}
		context.JSON(http.StatusOK, response)
	}
}

func (handler *TodoHandler) HandleCreateTodo() gin.HandlerFunc {
	return func(context *gin.Context) {
		claimsRaw, exists := context.Get("claims")
		if !exists {
			context.Error(errors.New("Thông tin xác thực bị mất sau khi giải mã"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		claims := claimsRaw.(*utils.Claims)
		requestRaw, exists := context.Get("request_data")
		if !exists {
			context.Error(errors.New("Dữ liệu request bị mất sau khi validate"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, parseResult := requestRaw.(*request.CreateTodoRequest)
		if !parseResult {
			context.Error(errors.New("Dữ liệu request không phải là CreateTodoRequest"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, error := handler.todoService.CreateTodo(claims.UserId, request)
		if error != nil {
			context.Error(error.Error)
			context.AbortWithStatus(error.StatusCode)
			return
		}
		context.JSON(http.StatusCreated, response)
	}
}

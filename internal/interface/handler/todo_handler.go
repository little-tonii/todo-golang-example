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

// GetTodoById 	godoc
// @Summary 	Tìm todo bằng id
// @Produce 	application/json
// @Tags 		Todo
// @Security	BearerAuth
// @Param 		id path int true "Request Param"
// @Success 	200 {object} response.GetTodoByIdResponse
// @Failure 	400 {object} godoc.ErrorsResponse
// @Failure		401 {object} godoc.ErrorResponse
// @Failure		404 {object} godoc.ErrorResponse
// @Failure		500 {object} godoc.ErrorResponse
// @Router		/todo/{id} [get]
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

// CreateTodo 	godoc
// @Summary 	Tạo mới todo
// @Produce 	application/json
// @Tags 		Todo
// @Security	BearerAuth
// @Param 		request body request.CreateTodoRequest true "Request Body"
// @Success 	200 {object} response.GetTodoByIdResponse
// @Failure 	400 {object} godoc.ErrorsResponse
// @Failure		401 {object} godoc.ErrorResponse
// @Failure		500 {object} godoc.ErrorResponse
// @Router		/todo/create [post]
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

// DeleteTodoById 	godoc
// @Summary 		Xóa todo theo id
// @Produce 		application/json
// @Tags 			Todo
// @Security		BearerAuth
// @Param 			id path int true "Request Param"
// @Success 		204
// @Failure 		400 {object} godoc.ErrorsResponse
// @Failure			401 {object} godoc.ErrorResponse
// @Failure 		404 {object} godoc.ErrorResponse
// @Failure			500 {object} godoc.ErrorResponse
// @Router			/todo/{id} [delete]
func (handler *TodoHandler) HandlerDeleteTodoById() gin.HandlerFunc {
	return func(context *gin.Context) {
		id, error := strconv.ParseInt(context.Param("id"), 10, 64)
		if error != nil || id <= 0 {
			context.Error(errors.New("Id phải là một số nguyên dương"))
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		appError := handler.todoService.DeleteTodoById(id)
		if appError != nil {
			context.Error(appError.Error)
			context.AbortWithStatus(appError.StatusCode)
			return
		}
		context.JSON(http.StatusNoContent, nil)
	}
}

// GetTodoList 	godoc
// @Summary 	Phân trang danh sách todo
// @Produce 	application/json
// @Tags 		Todo
// @Security	BearerAuth
// @Param 		page query int true "Query Param"
// @Param 		size query int true "Query Param"
// @Success 	200 {object} response.GetTodoListResponse
// @Failure 	400 {object} godoc.ErrorsResponse
// @Failure		401 {object} godoc.ErrorResponse
// @Failure		500 {object} godoc.ErrorResponse
// @Router		/todo/list [get]
func (handler *TodoHandler) HandleGetTodoList() gin.HandlerFunc {
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
		request, parseResult := requestRaw.(*request.GetTodoListRequest)
		if !parseResult {
			context.Error(errors.New("Dữ liệu request không phải là GetTodoListRequest"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, error := handler.todoService.GetTodoList(claims.UserId, request)
		if error != nil {
			context.Error(error.Error)
			context.AbortWithStatus(error.StatusCode)
			return
		}
		context.JSON(http.StatusOK, response)
	}
}

// UpdateTodoById 	godoc
// @Summary 		Cập nhật todo theo id
// @Produce 		application/json
// @Tags 			Todo
// @Security		BearerAuth
// @Param 			id path int true "Request Param"
// @Param			request body request.UpdateTodoByIdRequest true "Request Body"
// @Success 		200 {object} response.UpdateTodoByIdResponse
// @Failure 		400 {object} godoc.ErrorsResponse
// @Failure			401 {object} godoc.ErrorResponse
// @Failure 		404 {object} godoc.ErrorResponse
// @Failure			500 {object} godoc.ErrorResponse
// @Router			/todo/update/{id} [put]
func (handler *TodoHandler) HandleUpdateTodoById() gin.HandlerFunc {
	return func(context *gin.Context) {
		todoId, error := strconv.ParseInt(context.Param("id"), 10, 64)
		if error != nil || todoId <= 0 {
			context.Error(errors.New("Id phải là một số nguyên dương"))
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		requestRaw, exists := context.Get("request_data")
		if !exists {
			context.Error(errors.New("Dữ liệu request bị mất sau khi validate"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, parseResult := requestRaw.(*request.UpdateTodoByIdRequest)
		if !parseResult {
			context.Error(errors.New("Dữ liệu request không phải là UpdateTodoByIdRequest"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, appError := handler.todoService.UpdateTodoById(todoId, request)
		if appError != nil {
			context.Error(appError.Error)
			context.AbortWithStatus(appError.StatusCode)
			return
		}
		context.JSON(http.StatusOK, response)
	}
}

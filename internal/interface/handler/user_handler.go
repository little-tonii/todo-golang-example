package handler

import (
	"errors"
	"net/http"
	"todo-golang-example/internal/application/request"
	"todo-golang-example/internal/application/service"
	"todo-golang-example/pkg/utils"

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

// Register 	godoc
// @Summary 	Đăng ký
// @Produce 	application/json
// @Tags 		User
// @Param 		request body request.RegisterUserRequest true "Request Body"
// @Success 	201 {object} godoc.ErrorResponse
// @Failure		400 {object} godoc.ErrorsResponse
// @Failure		409 {object} godoc.ErrorResponse
// @Failure		500 {object} godoc.ErrorResponse
// @Router		/user/register [post]
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
			context.Error(error.Error)
			context.AbortWithStatus(error.StatusCode)
			return
		}
		context.JSON(http.StatusCreated, gin.H{"message": "Đăng ký thành công"})
	}
}

// Login 		godoc
// @Summary 	Đăng nhập
// @Produce 	application/json
// @Tags 		User
// @Param 		request body request.LoginUserRequest true "Request Body"
// @Success 	200 {object} response.LoginUserResponse
// @Failure		400 {object} godoc.ErrorsResponse
// @Failure		401 {object} godoc.ErrorResponse
// @Failure		500 {object} godoc.ErrorResponse
// @Router		/user/login [post]
func (handler *UserHandler) HandleLogin() gin.HandlerFunc {
	return func(context *gin.Context) {
		requestRaw, exists := context.Get("request_data")
		if !exists {
			context.Error(errors.New("Dữ liệu request bị mất sau khi validate"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		request, parseResult := requestRaw.(*request.LoginUserRequest)
		if !parseResult {
			context.Error(errors.New("Dữ liệu request không phải là LoginUserRequest"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		response, error := handler.userService.Login(request)
		if error != nil {
			context.Error(error.Error)
			context.AbortWithStatus(error.StatusCode)
			return
		}
		context.JSON(http.StatusOK, response)
	}
}

// Info 		godoc
// @Summary 	Thông tin người dùng
// @Produce 	application/json
// @Tags 		User
// @Security	BearerAuth
// @Success 	200 {object} response.GetUserInfoResponse
// @Failure		401 {object} godoc.ErrorResponse
// @Failure		500 {object} godoc.ErrorResponse
// @Router		/user/info [get]
func (handler *UserHandler) HandleInfo() gin.HandlerFunc {
	return func(context *gin.Context) {
		claimsRaw, exists := context.Get("claims")
		if !exists {
			context.Error(errors.New("Thông tin xác thực bị mất sau khi giải mã"))
			context.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		claims := claimsRaw.(*utils.Claims)
		response, error := handler.userService.Info(&claims.UserId)
		if error != nil {
			context.Error(error.Error)
			context.AbortWithStatus(error.StatusCode)
			return
		}
		context.JSON(http.StatusOK, response)
	}
}

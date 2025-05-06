package middleware

import (
	"errors"
	"net/http"
	"todo-golang-example/internal/infrastructure/config"
	"todo-golang-example/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		if token == "" {
			context.Error(errors.New("Người dùng chưa đăng nhập"))
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userId, err := utils.VerifyToken(config.Environment.JWT_SECRET_KEY, token)

		if err != nil {
			context.Error(errors.New("Người dùng chưa đăng nhập"))
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		context.Set("userId", userId)
		context.Next()
	}
}

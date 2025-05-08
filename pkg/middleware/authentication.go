package middleware

import (
	"errors"
	"net/http"
	"todo-golang-example/internal/shared/config"
	"todo-golang-example/pkg/utils"

	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {
			context.Error(errors.New("Người dùng chưa đăng nhập"))
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		const prefix = "Bearer "
		if len(authHeader) <= len(prefix) || authHeader[:len(prefix)] != prefix {
			context.Error(errors.New("Authorization header không hợp lệ"))
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		token := authHeader[len(prefix):]

		claims, err := utils.VerifyToken(config.Environment.JWT_SECRET_KEY, token)

		if err != nil {
			context.Error(errors.New(err.Error()))
			context.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		context.Set("claims", claims)
		context.Next()
	}
}

package middleware

import (
	"errors"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func BindingValidator[T any]() gin.HandlerFunc {
	var t T
	if tType := reflect.TypeOf(t); tType.Kind() != reflect.Struct {
		panic("BindingValidator chỉ chấp nhận kiểu struct")
	}
	return func(context *gin.Context) {
		var requestData T
		if error := context.ShouldBindJSON(&requestData); error != nil {
			var validatorError validator.ValidationErrors
			if errors.As(error, &validatorError) {
				for _, error := range validatorError {
					field := error.Field()
					tag := error.Tag()
					switch tag {
					default:
						context.Error(errors.New("Lỗi không không xác định khi validation"))
					}
				}
				context.AbortWithStatus(http.StatusBadRequest)
				return
			}
			context.Error(errors.New("Định dạng json không hợp lệ"))
			context.AbortWithStatus(http.StatusBadRequest)
			return
		}
		context.Set("request_data", requestData)
		context.Next()
	}
}

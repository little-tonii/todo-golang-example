package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessToken(secretKey string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func GenerateRefreshToken(secretKey string, userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  userId,
		"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}

func VerifyToken(secretKey string, token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("Có lỗi xảy ra")
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		return "", errors.New("Không thể phân tích token")
	}

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid {
		return "", errors.New("Token không hợp lệ")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return "", errors.New("Token claims không hợp lệ")
	}
	userId := claims["id"].(string)
	return userId, nil
}

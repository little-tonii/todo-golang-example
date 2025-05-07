package service

import (
	"errors"
	"net/http"
	"todo-golang-example/internal/application/request"
	"todo-golang-example/internal/application/response"
	"todo-golang-example/internal/domain/entity"
	"todo-golang-example/internal/domain/repository"
	"todo-golang-example/internal/shared/common"
	"todo-golang-example/pkg/utils"

	"gorm.io/gorm"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) Register(request *request.RegisterUserRequest) *common.ApplicationError {
	hashedPassword, error := utils.HashPassword(request.Password)
	if error != nil {
		return common.NewApplicationError(
			http.StatusInternalServerError,
			errors.New("Có lỗi trong quá trình mã hóa mật khẩu"),
		)
	}
	_, error = userService.userRepository.GetByEmail(request.Email)
	if error != nil && !errors.Is(error, gorm.ErrRecordNotFound) {
		return common.NewApplicationError(http.StatusInternalServerError, error)
	}
	userEntity := &entity.UserEntity{
		Email:          request.Email,
		HashedPassword: hashedPassword,
	}
	error = userService.userRepository.Create(userEntity)
	if error != nil {
		return common.NewApplicationError(
			http.StatusConflict,
			errors.New("Email đã được sử dụng"),
		)
	}
	return nil
}

func (userService *UserService) Login(request *request.LoginUserRequest) (*response.LoginUserResponse, error) {
	return nil, nil
}

func (userService *UserService) Info(userId *int64) (*response.GetUserInfoResponse, error) {
	return nil, nil
}

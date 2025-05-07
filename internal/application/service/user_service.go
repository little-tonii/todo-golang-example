package service

import (
	"todo-golang-example/internal/application/request"
	"todo-golang-example/internal/application/response"
	"todo-golang-example/internal/domain/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (userService *UserService) Register(request *request.RegisterUserRequest) error {
	return nil
}

func (userService *UserService) Login(request *request.LoginUserRequest) (*response.LoginUserResponse, error) {
	return nil, nil
}

func (userService *UserService) Info(userId *int64) (*response.GetUserInfoResponse, error) {
	return nil, nil
}

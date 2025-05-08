package service

import (
	"errors"
	"fmt"
	"net/http"
	"todo-golang-example/internal/application/request"
	"todo-golang-example/internal/application/response"
	"todo-golang-example/internal/domain/entity"
	"todo-golang-example/internal/domain/repository"
	"todo-golang-example/internal/shared/common"
	"todo-golang-example/internal/shared/config"
	pkgUtils "todo-golang-example/pkg/utils"

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
	hashedPassword, error := pkgUtils.HashPassword(request.Password)
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

func (userService *UserService) Login(request *request.LoginUserRequest) (*response.LoginUserResponse, *common.ApplicationError) {
	userEntity, error := userService.userRepository.GetByEmail(request.Email)
	if error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return nil, common.NewApplicationError(
				http.StatusUnauthorized,
				errors.New(fmt.Sprintf("Email %s chưa đăng ký tài khoản", request.Email)),
			)
		} else {
			return nil, common.NewApplicationError(http.StatusInternalServerError, error)
		}
	}
	if !pkgUtils.CheckPasswordHash(request.Password, userEntity.HashedPassword) {
		return nil, common.NewApplicationError(
			http.StatusUnauthorized,
			errors.New("Tài khoản hoặc mật khẩu không chính xác"),
		)
	}
	refreshToken, error := pkgUtils.GenerateRefreshToken(config.Environment.JWT_SECRET_KEY, userEntity.Id)
	if error != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, error)
	}
	accessToken, error := pkgUtils.GenerateAccessToken(config.Environment.JWT_SECRET_KEY, userEntity.Id)
	userEntity.RefreshToken = refreshToken
	error = userService.userRepository.Update(userEntity)
	if error != nil {
		return nil, common.NewApplicationError(http.StatusInternalServerError, error)
	}
	return &response.LoginUserResponse{
		RefreshToken: refreshToken,
		AccessToken:  accessToken,
	}, nil
}

func (userService *UserService) Info(userId *int64) (*response.GetUserInfoResponse, *common.ApplicationError) {
	userEntity, error := userService.userRepository.GetById(*userId)
	if error != nil {
		if errors.Is(error, gorm.ErrRecordNotFound) {
			return nil, common.NewApplicationError(
				http.StatusUnauthorized,
				errors.New(fmt.Sprintf("Người dùng không tồn tại")),
			)
		} else {
			return nil, common.NewApplicationError(http.StatusInternalServerError, error)
		}
	}
	return &response.GetUserInfoResponse{Id: userEntity.Id, Email: userEntity.Email}, nil
}

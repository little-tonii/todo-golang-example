package repositoryimpl

import (
	"errors"
	"todo-microservice/internal/domain/entity"
	"todo-microservice/internal/infrastructure/model"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	database *gorm.DB
}

func NewUserRepositoryImpl(database *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		database: database,
	}
}

func (userRepository *UserRepositoryImpl) CreateUser(userEntity *entity.UserEntity) error {
	userModel := model.UserModel{
		Email:          userEntity.Email,
		HashedPassword: userEntity.HashedPassword,
	}
	result := userRepository.database.Create(&userModel)
	if result.Error != nil {
		return errors.New("Email đã được sử dụng")
	}
	return nil
}

func (userRepository *UserRepositoryImpl) GetUserByEmail(email string) (*entity.UserEntity, error) {
	userModel := model.UserModel{Email: email}
	result := userRepository.database.First(&userModel)
	if result.Error != nil {
		return nil, errors.New("Email không tồn tại")
	}
	return userModel.ToEntity(), nil
}

func (userRepository *UserRepositoryImpl) GetUserById(id int64) (*entity.UserEntity, error) {
	userModel := model.UserModel{Id: id}
	result := userRepository.database.First(&userModel)
	if result.Error != nil {
		return nil, errors.New("Id không tồn tại")
	}
	return userModel.ToEntity(), nil
}

package repositoryimpl

import (
	"errors"
	"todo-golang-example/internal/domain/entity"
	"todo-golang-example/internal/infrastructure/model"

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

func (userRepository *UserRepositoryImpl) Create(userEntity *entity.UserEntity) error {
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

func (userRepository *UserRepositoryImpl) GetById(email string) (*entity.UserEntity, error) {
	userModel := model.UserModel{Email: email}
	result := userRepository.database.First(&userModel)
	if result.Error != nil {
		return nil, errors.New("Email không tồn tại")
	}
	return userModel.ToEntity(), nil
}

func (userRepository *UserRepositoryImpl) GetByEmail(id int64) (*entity.UserEntity, error) {
	userModel := model.UserModel{Id: id}
	result := userRepository.database.First(&userModel)
	if result.Error != nil {
		return nil, errors.New("Id không tồn tại")
	}
	return userModel.ToEntity(), nil
}

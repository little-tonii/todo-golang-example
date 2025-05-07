package repositoryimpl

import (
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
	return userRepository.database.Create(&userModel).Error
}

func (userRepository *UserRepositoryImpl) GetByEmail(email string) (*entity.UserEntity, error) {
	userModel := model.UserModel{Email: email}
	result := userRepository.database.First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}

func (userRepository *UserRepositoryImpl) GetById(id int64) (*entity.UserEntity, error) {
	userModel := model.UserModel{Id: id}
	result := userRepository.database.First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}

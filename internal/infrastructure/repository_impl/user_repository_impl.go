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
	var userModel model.UserModel
	result := userRepository.database.
		Where(&model.UserModel{Email: email}).
		First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}

func (userRepository *UserRepositoryImpl) GetById(id int64) (*entity.UserEntity, error) {
	var userModel model.UserModel
	result := userRepository.database.
		Where(&model.UserModel{Id: id}).
		First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}

func (userRepository *UserRepositoryImpl) Update(userEntity *entity.UserEntity) error {
	userModel := model.UserModel{
		HashedPassword: userEntity.HashedPassword,
		RefreshToken:   userEntity.RefreshToken,
	}
	result := userRepository.database.
		Where(&model.UserModel{Id: userEntity.Id}).
		Updates(&userModel)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

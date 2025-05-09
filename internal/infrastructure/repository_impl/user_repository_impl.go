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

func (repository *UserRepositoryImpl) Create(userEntity *entity.UserEntity) error {
	userModel := model.UserModel{
		Email:          userEntity.Email,
		HashedPassword: userEntity.HashedPassword,
	}
	error := repository.database.Create(&userModel).
		Error
	if error != nil {
		return error
	}
	userEntity.Id = userModel.Id
	return nil
}

func (repository *UserRepositoryImpl) FindByEmail(email string) (*entity.UserEntity, error) {
	var userModel model.UserModel
	result := repository.database.
		Where("email = ?", email).
		First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}

func (repository *UserRepositoryImpl) FindById(id int64) (*entity.UserEntity, error) {
	var userModel model.UserModel
	result := repository.database.
		Where("id = ?", id).
		First(&userModel)
	if result.Error != nil {
		return nil, result.Error
	}
	return userModel.ToEntity(), nil
}

func (repository *UserRepositoryImpl) Update(userEntity *entity.UserEntity) error {
	result := repository.database.
		Model(&model.UserModel{}).
		Where("id = ?", userEntity.Id).
		Updates(map[string]any{
			"hashed_password": userEntity.HashedPassword,
			"refresh_token":   userEntity.RefreshToken,
		})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

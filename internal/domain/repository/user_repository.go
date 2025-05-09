package repository

import (
	"todo-golang-example/internal/domain/entity"
)

type UserRepository interface {
	Create(user *entity.UserEntity) error
	FindById(id int64) (*entity.UserEntity, error)
	FindByEmail(email string) (*entity.UserEntity, error)
	Update(user *entity.UserEntity) error
}

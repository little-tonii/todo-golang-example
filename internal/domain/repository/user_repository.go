package repository

import (
	"todo-golang-example/internal/domain/entity"
)

type UserRepository interface {
	Create(user *entity.UserEntity) error
	GetById(id int64) (*entity.UserEntity, error)
	GetByEmail(email string) (*entity.UserEntity, error)
	Update(user *entity.UserEntity) error
}

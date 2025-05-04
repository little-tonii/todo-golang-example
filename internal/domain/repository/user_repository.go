package repository

import "todo-microservice/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.UserEntity) error
	GetById(id int64) (*entity.UserEntity, error)
	GetByEmail(email string) (*entity.UserEntity, error)
}

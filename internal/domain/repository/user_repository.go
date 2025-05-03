package repository

import "todo-microservice/internal/domain/entity"

type UserRepository interface {
	CreateUser(user *entity.UserEntity) error
	GetUserById(id int64) (*entity.UserEntity, error)
	GetUserByEmail(email string) (*entity.UserEntity, error)
}

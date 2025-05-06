package repository

import (
	"sync"
	"todo-golang-example/internal/domain/entity"
	"todo-golang-example/internal/infrastructure/config"
	repositoryimpl "todo-golang-example/internal/infrastructure/repository_impl"
)

type UserRepository interface {
	Create(user *entity.UserEntity) error
	GetById(id int64) (*entity.UserEntity, error)
	GetByEmail(email string) (*entity.UserEntity, error)
}

var (
	userRepository UserRepository
	once           sync.Once
)

func GetUserRepository() UserRepository {
	once.Do(func() {
		userRepository = repositoryimpl.NewUserRepositoryImpl(config.GetDatabase())
	})
	return userRepository
}

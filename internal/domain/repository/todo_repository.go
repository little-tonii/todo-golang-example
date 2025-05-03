package repository

import "todo-microservice/internal/domain/entity"

type TodoRepository interface {
	Create(todo *entity.TodoEntity) error
	Update(todo *entity.TodoEntity) error
	DeleteById(id int64) error
	GetById(id int64) (*entity.TodoEntity, error)
	GetAll() ([]*entity.TodoEntity, error)
}

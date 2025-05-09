package repository

import "todo-golang-example/internal/domain/entity"

type TodoRepository interface {
	Create(todo *entity.TodoEntity) error
	Update(todo *entity.TodoEntity) error
	DeleteById(id int64) error
	FindById(id int64) (*entity.TodoEntity, error)
	List(page int64, size int64) ([]*entity.TodoEntity, error)
}

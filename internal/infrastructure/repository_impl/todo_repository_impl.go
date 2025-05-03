package repositoryimpl

import (
	"errors"
	"todo-microservice/internal/domain/entity"
	"todo-microservice/internal/infrastructure/model"

	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	database *gorm.DB
}

func NewTodoRepositoryImpl(database *gorm.DB) *TodoRepositoryImpl {
	return &TodoRepositoryImpl{
		database: database,
	}
}

func (todoRepositoryImpl *TodoRepositoryImpl) createTodo(todoEntity *entity.TodoEntity) error {
	todoModel := model.TodoModel{
		Description: todoEntity.Description,
		Title:       todoEntity.Title,
	}
	result := todoRepositoryImpl.database.Create(&todoModel)
	if result.Error != nil {
		return errors.New("Có lỗi xảy ra khi tạo mới todo")
	}
	return nil
}

func (todoRepositoryImpl *TodoRepositoryImpl) updateTodo(todoEntity *entity.TodoEntity) error {
	todoModel := model.TodoModel{
		Description: todoEntity.Description,
		Title:       todoEntity.Title,
	}
	result := todoRepositoryImpl.database.Save(&todoModel)
	if result.Error != nil {
		return errors.New("Có lỗi xảy ra khi cập nhật todo")
	}
	return nil
}

func (todoRepositoryImpl *TodoRepositoryImpl) deleteTodoById(todoEntity *entity.TodoEntity) error {
	todoModel := model.TodoModel{
		Id: todoEntity.Id,
	}
	result := todoRepositoryImpl.database.Delete(&todoModel)
	if result.Error != nil {
		return errors.New("Có lỗi xảy ra khi xóa todo")
	}
	return nil
}

func (todoRepositoryImpl *TodoRepositoryImpl) getTodoById

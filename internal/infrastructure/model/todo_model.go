package model

import (
	"time"
	"todo-golang-example/internal/domain/entity"
)

type TodoModel struct {
	Id          int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Title       string    `gorm:"column:title;not null"`
	Description string    `gorm:"column:description;not null"`
	UserId      int64     `gorm:"column:user_id;not null;index"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;autoUpdateTime"`
}

func (TodoModel) TableName() string {
	return "todos"
}

func (todoModel *TodoModel) ToEntity() *entity.TodoEntity {
	return &entity.TodoEntity{
		Id:          todoModel.Id,
		Title:       todoModel.Title,
		Description: todoModel.Description,
		UserId:      todoModel.UserId,
		CreatedAt:   todoModel.CreatedAt,
		UpdatedAt:   todoModel.UpdatedAt,
	}
}

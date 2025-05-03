package model

import (
	"time"
	"todo-microservice/internal/domain/entity"
)

type UserModel struct {
	Id             int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Email          string    `gorm:"column:email;unique;not null"`
	HashedPassword string    `gorm:"column:hashed_password;not null"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;autoUpdateTime"`
}

func (userModel *UserModel) ToEntity() *entity.UserEntity {
	return &entity.UserEntity{
		Id:             userModel.Id,
		Email:          userModel.Email,
		HashedPassword: userModel.HashedPassword,
		CreatedAt:      userModel.CreatedAt,
		UpdatedAt:      userModel.UpdatedAt,
	}
}

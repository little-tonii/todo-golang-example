package model

import (
	"time"
	"todo-golang-example/internal/domain/entity"
)

type UserModel struct {
	Id             int64     `gorm:"column:id;primaryKey;autoIncrement"`
	Email          string    `gorm:"column:email;unique;not null"`
	HashedPassword string    `gorm:"column:hashed_password;not null"`
	RefreshToken   string    `gorm:"column:refresh_token;not null"`
	CreatedAt      time.Time `gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;not null;autoUpdateTime"`
}

func (UserModel) TableName() string {
	return "users"
}

func (userModel *UserModel) ToEntity() *entity.UserEntity {
	return &entity.UserEntity{
		Id:             userModel.Id,
		Email:          userModel.Email,
		HashedPassword: userModel.HashedPassword,
		RefreshToken:   userModel.RefreshToken,
		CreatedAt:      userModel.CreatedAt,
		UpdatedAt:      userModel.UpdatedAt,
	}
}

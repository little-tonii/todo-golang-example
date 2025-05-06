package entity

import "time"

type UserEntity struct {
	Id             int64
	Email          string
	HashedPassword string
	RefreshToken   string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

package entity

import "time"

type TodoEntity struct {
	Id          int64
	Title       string
	Description string
	UserId      int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

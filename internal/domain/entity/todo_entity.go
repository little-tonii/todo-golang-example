package entity

import "time"

type TodoEntity struct {
	Id          int64
	Title       string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

package entity

import "time"

type Posts struct {
	ID        int32
	Title     string `validate:"required,min=5"`
	Body      string `validate:"required,min=50"`
	CreatedAt time.Time
}

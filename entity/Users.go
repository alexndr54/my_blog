package entity

import (
	"time"
)

type Users struct {
	ID       int32
	Nama     string `validate:"required,min=5,alpha"`
	Email    string `validate:"required,email,min=5"`
	Password string `validate:"required,min=8"`
	Created  time.Time
}

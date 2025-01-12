package repository

import (
	"blog/entity"
	"context"
)

type RepositoryUsers interface {
	AddUsers(ctx context.Context, users *entity.Users) (*entity.Users, error)
	DeleteUsers(ctx context.Context, users *entity.Users) (*entity.Users, error)
	UpdateUsers(ctx context.Context, users *entity.Users) (*entity.Users, error)
	FindUsersByEmail(ctx context.Context, users *entity.Users) (*entity.Users, error)
}

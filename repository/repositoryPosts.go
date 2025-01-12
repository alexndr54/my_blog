package repository

import (
	"blog/entity"
	"context"
)

type repositoryPosts interface {
	AddPosts(ctx context.Context, posts entity.Posts) error
	DeletePosts(ctx context.Context, posts entity.Posts) error
	GetPosts(ctx context.Context, posts entity.Posts) (entity.Posts, error)
	GetAllPosts(ctx context.Context) ([]entity.Posts, error)
}

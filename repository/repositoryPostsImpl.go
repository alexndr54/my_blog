package repository

import (
	"blog/entity"
	"context"
	"database/sql"
	"errors"
)

type PostsRepositoryImpl struct {
	DB *sql.DB
}

func (p PostsRepositoryImpl) AddPosts(ctx context.Context, posts entity.Posts) error {
	query := "INSERT INTO posts (title,body) VALUES (?,?)"
	_, err := p.DB.ExecContext(ctx, query, posts.Title, posts.Body)
	if err != nil {
		return err
	}
	return nil
}

func (p PostsRepositoryImpl) DeletePosts(ctx context.Context, posts entity.Posts) error {
	if posts.ID == 0 {
		return errors.New("posts id is zero")
	}

	query := "DELETE FROM posts WHERE id = ?"
	result, err := p.DB.ExecContext(ctx, query, posts.ID)
	if err != nil {
		return err
	}

	affected, _ := result.RowsAffected()
	if affected == 0 {
		return errors.New("posts affected is zero")
	}

	return nil
}

func (p PostsRepositoryImpl) GetPosts(ctx context.Context, posts entity.Posts) (entity.Posts, error) {
	if posts.ID == 0 {
		return entity.Posts{}, errors.New("posts id is zero")
	}

	query := "SELECT id, title, body,created_at FROM posts WHERE id = ?"
	rows, err := p.DB.QueryContext(ctx, query, posts.ID)
	if err != nil {
		return entity.Posts{}, err
	}

	defer rows.Close()
	var post entity.Posts
	for rows.Next() {
		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt); err != nil {
			return entity.Posts{}, err
		}

	}

	return post, nil
}

func (p PostsRepositoryImpl) GetAllPosts(ctx context.Context) ([]entity.Posts, error) {
	query := "SELECT id, title, body,created_at FROM posts"
	rows, err := p.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []entity.Posts
	for rows.Next() {
		var post entity.Posts
		if err := rows.Scan(&post.ID, &post.Title, &post.Body, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

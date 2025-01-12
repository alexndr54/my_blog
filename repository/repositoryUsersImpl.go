package repository

import (
	"blog/entity"
	"context"
	"database/sql"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type RepositoryUsersImpl struct {
	DB *sql.DB
}

func (r RepositoryUsersImpl) AddUsers(ctx context.Context, users *entity.Users) (*entity.Users, error) {
	query := "INSERT INTO users (nama,email,password) VALUES (?,?,?)"
	hash, err := bcrypt.GenerateFromPassword([]byte(users.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	result, err := r.DB.ExecContext(ctx, query, users.Nama, users.Email, hash)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	users.ID = int32(id)
	return users, nil
}

func (r RepositoryUsersImpl) DeleteUsers(ctx context.Context, users *entity.Users) (*entity.Users, error) {
	result, err := r.DB.ExecContext(ctx, "DELETE FROM users WHERE id=?", users.ID)
	if err != nil {
		return nil, err
	}

	_, err = result.RowsAffected()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r RepositoryUsersImpl) UpdateUsers(ctx context.Context, users *entity.Users) (*entity.Users, error) {
	query := "UPDATE users SET nama=?, email=? WHERE id=?"
	res, err := r.DB.ExecContext(ctx, query, users.Nama, users.Email, users.ID)
	if err != nil {
		return nil, err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r RepositoryUsersImpl) FindUsersByEmail(ctx context.Context, users *entity.Users) (*entity.Users, error) {
	query := "SELECT id,nama,password, email FROM users WHERE email=?"
	res, err := r.DB.QueryContext(ctx, query, users.Email)
	if err != nil {
		return nil, err
	}

	if res.Next() {
		res.Scan(&users.ID, &users.Nama, &users.Password, &users.Email)
	} else {
		return users, errors.New("user not found")
	}

	return users, nil
}

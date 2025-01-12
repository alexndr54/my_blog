package main

import (
	"blog/config"
	"blog/entity"
	"blog/repository"
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"testing"
)

func TestAddUsers(t *testing.T) {
	ctx := context.Background()
	db := config.GetConnection()
	password := []byte("123")
	repo := repository.RepositoryUsersImpl{db}
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	assert.Nil(t, err)

	var users *entity.Users = &entity.Users{
		Nama:     "Alex",
		Email:    "alex@gmail.com",
		Password: string(hash),
	}
	users, err = repo.AddUsers(ctx, users)
	assert.Nil(t, err)
	t.Log(users)
	id := users.ID
	t.Log("Berhasil ditambahkan dengan ID: " + strconv.Itoa(int(users.ID)))

	t.Run("update", func(t *testing.T) {
		users.ID = id
		users.Nama = "Ini sudah diupdate"
		users.Email = "ini email sudah di update"

		user, err := repo.UpdateUsers(ctx, users)
		assert.Nil(t, err)
		t.Log("Berhasil diupdate menjadi: ", user)
	})

	t.Run("select", func(t *testing.T) {
		users.Email = "alex@gmail.com"
		get, err := repo.FindUsersByEmail(ctx, users)
		assert.Nil(t, err)
		t.Log("Berhasil get users: ", get)
	})

	t.Run("delete", func(t *testing.T) {
		users.ID = id
		del, err := repo.DeleteUsers(ctx, users)
		assert.Nil(t, err)
		fmt.Println("Berhasil dihapus user: ", del)
	})

}

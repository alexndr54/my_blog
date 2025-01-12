package main

import (
	"blog/config"
	"blog/entity"
	"blog/helper"
	"blog/repository"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestLogin(t *testing.T) {
	app := fiber.New(fiber.Config{})
	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	Validate, Trans := repository.SetLanguageID()

	formInput := entity.Users{
		Nama:     "Login",
		Email:    "sialexsofficial@gmail.com",
		Password: "sialexsofficial@gmail.com",
	}

	Validation := Validate.Struct(formInput)

	err, _ := repository.ValidasiResponse(Validate, Trans, Validation)
	if err == nil {
		repo := repository.RepositoryUsersImpl{config.GetConnection()}
		users, FindErr := repo.FindUsersByEmail(context.Background(), &formInput)
		t.Log("FindUsersByEmail: ", users)

		if FindErr != nil {
			t.Log("User tidak ditemukan")
			t.FailNow()
		}

		checkLogin := helper.VerifyPassword(users.Password, formInput.Password)
		t.Log("Login: ", checkLogin)
		t.Log("users.Password: ", users.Password)
		t.Log("formInput.Password: ", "sialexsofficial@gmail.com")
		t.FailNow()

	}

	t.Log("Validasi gagal dilewati")
	t.FailNow()

}

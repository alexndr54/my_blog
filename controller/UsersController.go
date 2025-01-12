package controller

import (
	"blog/config"
	"blog/entity"
	"blog/helper"
	"blog/repository"
	"context"
	"github.com/gofiber/fiber/v2"
)

func NewRegister(c *fiber.Ctx) error {

	Validate, Trans := repository.SetLanguageID()

	formInput := entity.Users{
		Nama:     c.FormValue("nama"),
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}

	Validation := Validate.Struct(formInput)

	err, response := repository.ValidasiResponse(c, Trans, Validation)
	response.Validation.ValidationValue = map[string]string{
		"Nama":  formInput.Nama,
		"Email": formInput.Email,
	}
	if err == nil {
		repo := repository.RepositoryUsersImpl{config.GetConnection()}
		_, err := repo.AddUsers(context.Background(), &formInput)
		if err == nil {
			response.Validation = entity.ValidationResponse{
				Status: entity.StatusResponse{
					Label:   repository.Success,
					Message: repository.RegisterSuccess,
				},
			}
		} else {
			response.Validation = entity.ValidationResponse{
				Status: entity.StatusResponse{
					Label:   repository.Danger,
					Message: repository.RegisterFailedHaveAccount,
				},
			}
		}

	}

	return c.Render("auth/register", response, "partials/main")

}

func NewLogin(c *fiber.Ctx) error {

	Validate, Trans := repository.SetLanguageID()

	formInput := entity.Users{
		Nama:     "Login",
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	formInput2 := formInput

	Validation := Validate.Struct(formInput)

	err, response := repository.ValidasiResponse(c, Trans, Validation)
	response.Validation.ValidationValue = map[string]string{
		"Email": formInput.Email,
	}

	if err == nil {
		repo := repository.RepositoryUsersImpl{config.GetConnection()}
		users, FindErr := repo.FindUsersByEmail(context.Background(), &formInput2)

		if FindErr != nil {
			response.Validation = entity.ValidationResponse{
				Status: entity.StatusResponse{
					Label:   repository.Danger,
					Message: repository.LoginFailedUserNotFound,
				},
			}
			return c.Render("auth/login", response, "partials/main")
		}

		checkLogin := helper.VerifyPassword(users.Password, formInput.Password)
		if checkLogin == false {
			response.Validation = entity.ValidationResponse{
				Status: entity.StatusResponse{
					Label:   repository.Danger,
					Message: repository.LoginWrongPassword,
				},
			}
			return c.Render("auth/login", response, "partials/main")
		}

		err := config.SetSession(c, "email", c.FormValue("email"))
		if err != nil {
			response.Validation = entity.ValidationResponse{
				Status: entity.StatusResponse{
					Label:   repository.Danger,
					Message: repository.LoginSessionFailed,
				},
			}
			return c.Render("auth/login", response, "partials/main")
		}

		return c.Redirect("/app/blog/list")

	}

	return c.Render("auth/login", *response, "partials/main")

}

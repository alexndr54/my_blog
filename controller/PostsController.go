package controller

import (
	"blog/config"
	"blog/entity"
	"blog/helper"
	"blog/repository"
	"context"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func NewPosts(c *fiber.Ctx) error {
	Validate, Trans := repository.SetLanguageID()

	form := entity.Posts{
		Title: c.FormValue("title"),
		Body:  c.FormValue("body"),
	}

	Validation := Validate.Struct(form)

	err, response := repository.ValidasiResponse(c, Trans, Validation)
	response.Validation.ValidationValue = map[string]string{
		"Title": form.Title,
		"Body":  form.Body,
	}

	if err == nil {
		repo := repository.PostsRepositoryImpl{config.GetConnection()}
		err := repo.AddPosts(context.Background(), form)
		if err != nil {
			response.Validation.Status = entity.StatusResponse{
				Label:   repository.Danger,
				Message: repository.AddPostsFailedUnknown,
			}
		} else {
			response.Validation.Status = entity.StatusResponse{
				Label:   repository.Success,
				Message: repository.AddPostsSuccess,
			}
		}
	}

	return c.Render("app/blog/add", response, "partials/main")

}

func ListPosts(c *fiber.Ctx) error {
	repo := repository.PostsRepositoryImpl{config.GetConnection()}
	posts, _ := repo.GetAllPosts(context.Background())

	Data := entity.ResponseDataRoute{
		Title:    "Daftar Posts",
		IsLogin:  helper.IsLogin(c),
		Optional: posts,
	}

	return c.Render("app/blog/list", Data, "partials/main")

}

func DeletedPost(c *fiber.Ctx) error {
	repo := repository.PostsRepositoryImpl{config.GetConnection()}
	id, _ := strconv.Atoi(c.Params("id"))
	users := entity.Posts{ID: int32(id)}
	delete := repo.DeletePosts(context.Background(), users)
	if delete != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(200)
}

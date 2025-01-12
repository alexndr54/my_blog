package main

import (
	"blog/Middleware"
	"blog/config"
	"blog/controller"
	"blog/entity"
	"blog/helper"
	"blog/repository"
	"context"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	_, app := config.GetFiberConfig()
	app.Get("/", func(c *fiber.Ctx) error {
		repo := repository.PostsRepositoryImpl{DB: config.GetConnection()}
		AllPost, _ := repo.GetAllPosts(context.Background())
		Data := entity.ResponseDataRoute{
			Title:    "Blog",
			IsLogin:  helper.IsLogin(c),
			Optional: AllPost,
		}
		return c.Render("index", Data, "partials/main")
	})
	app.Get("/read/:id", func(c *fiber.Ctx) error {
		repo := repository.PostsRepositoryImpl{DB: config.GetConnection()}
		id, _ := strconv.Atoi(c.Params("id"))
		post := entity.Posts{ID: int32(id)}
		Post, _ := repo.GetPosts(context.Background(), post)

		Data := entity.ResponseDataRoute{
			Title:    "Blog",
			IsLogin:  helper.IsLogin(c),
			Optional: Post,
		}
		return c.Render("read", Data, "partials/main")
	})

	app.Get("/logout", controller.Logout)

	Login := app.Group("/auth", Middleware.LoginNoAuthMiddleware)
	Login.Get("/login", func(c *fiber.Ctx) error {
		Data := entity.ResponseDataRoute{
			Title:   "Login",
			IsLogin: helper.IsLogin(c),
		}
		return c.Render("auth/login", Data, "partials/main")

	})
	Login.Get("/register", func(c *fiber.Ctx) error {
		Data := entity.ResponseDataRoute{
			Title:   "Register",
			IsLogin: helper.IsLogin(c),
		}

		return c.Render("auth/register", Data, "partials/main")

	})

	Login.Post("/register", controller.NewRegister).Name("NLogin")
	Login.Post("/login", controller.NewLogin).Name("NLogin")

	Dashboard := app.Group("/app", Middleware.AuthMiddleware)

	Blog := Dashboard.Group("/blog")
	Blog.Get("/add", func(c *fiber.Ctx) error {
		Data := entity.ResponseDataRoute{
			Title:   "Tambah Artikel",
			IsLogin: helper.IsLogin(c),
		}
		return c.Render("app/blog/add", Data, "partials/main")
	})
	Blog.Post("/add", controller.NewPosts)

	Blog.Get("/list", controller.ListPosts)
	Blog.Get("/delete/:id", controller.DeletedPost)

	log.Fatal(app.Listen(":2003"))
}

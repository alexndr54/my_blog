package Middleware

import (
	"blog/config"
	"blog/entity"
	"blog/repository"
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var (
	store = session.New()
)

func AuthMiddleware(c *fiber.Ctx) error {
	_, Sess := config.GetSession(c)
	email := Sess.Get("email")
	if email == nil {
		return c.Redirect("/auth/login")
	}
	users := entity.Users{Email: email.(string)}

	repo := repository.RepositoryUsersImpl{DB: config.GetConnection()}
	_, err := repo.FindUsersByEmail(context.Background(), &users)
	if err != nil {
		return c.Redirect("/auth/login")
	}

	c.Locals("email", users.Email)
	return c.Next()
}

func LoginNoAuthMiddleware(c *fiber.Ctx) error {
	_, Sess := config.GetSession(c)
	email := Sess.Get("email")
	if email != nil {
		return c.Redirect("/app/blog/list")
	}

	return c.Next()
}

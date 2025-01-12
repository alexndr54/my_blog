package controller

import (
	"blog/config"
	"github.com/gofiber/fiber/v2"
)

func Logout(c *fiber.Ctx) error {
	_, sess := config.GetSession(c)
	sess.Delete("email")
	sess.Destroy()
	sess.Save()
	c.Locals("email", nil)

	return c.Redirect("/")
}

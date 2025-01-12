package helper

import (
	"github.com/gofiber/fiber/v2"
)

func IsLogin(c *fiber.Ctx) bool {
	if c.Locals("email") == nil {
		return false
	}
	return c.Locals("email") != "" && len(c.Locals("email").(string)) > 5
}

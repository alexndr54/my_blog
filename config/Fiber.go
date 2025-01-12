package config

import (
	"blog/helper"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/valyala/fasthttp"
	"html/template"
)

func GetFiberConfig() (*fiber.Ctx, *fiber.App) {
	engine := html.New("./views", ".html")
	engine.AddFuncMap(template.FuncMap{
		"ShowMessage":       helper.ShowMessage,
		"ShowFailedMessage": helper.ShowFailedMessage,
		"ShowValueMessage":  helper.ShowValueMessage,
	})
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})

	return ctx, app
}

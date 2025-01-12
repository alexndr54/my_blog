package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"
	"testing"
)

func TestSession(t *testing.T) {
	store := session.New()
	app := fiber.New(fiber.Config{})

	ctx := app.AcquireCtx(&fasthttp.RequestCtx{})
	defer app.ReleaseCtx(ctx)

	sess, err := store.Get(ctx)
	assert.Nil(t, err)
	t.Log(sess)
	t.Log("Token: " + sess.ID())

	sess.Set("email", "world")
	t.Log("Email: ", sess.Get("email"))
}

func TestAddSession(t *testing.T) {

}

func TestCheckSession(t *testing.T) {

}

func TestAddCheckSession(t *testing.T) {
	TestAddSession(t)
	TestCheckSession(t)
}

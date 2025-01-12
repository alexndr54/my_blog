package config

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/redis/v3"
	"time"
)

func configSession() (*redis.Storage, *session.Store) {
	Storages := redis.New(redis.Config{
		Host: "127.0.0.1",
		Port: 6379,
	})
	store := session.New(session.Config{
		Expiration:     3 * time.Hour,
		CookieHTTPOnly: true,
		CookieSecure:   false,
		CookiePath:     "/",
		Storage:        Storages,
	})

	return Storages, store
}
func GetSession(c *fiber.Ctx) (*redis.Storage, *session.Session) {
	storages, store := configSession()

	StoreSession, err := store.Get(c)
	if err != nil {
		panic(err)
	}

	return storages, StoreSession
}
func SetSession(c *fiber.Ctx, key string, value string) error {
	_, store := configSession()
	StoreSession, err := store.Get(c)
	if err != nil {
		return err
	}

	StoreSession.Set(key, value)
	StoreSession.Save()

	return nil
}

package routes

import (
	"time"

	"github.com/bachiitter/api/pkg/unsplash"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func UnsplashRouter(app fiber.Router) {
	router := app.Group("/unsplash")

	router.Use(cache.New(cache.Config{
		Next: func(c *fiber.Ctx) bool {
			return c.Query("refresh") == "true"
		},
		Expiration:   30 * time.Minute,
		CacheControl: true,
	}))

	router.Get("/photos/:username", func(c *fiber.Ctx) error {
		username := c.Params("username")

		data, err := unsplash.GetUserPhotos(username)
		if err != nil {
			return err
		}

		return c.JSON(data)
	})
	router.Get("/stats/:username", func(c *fiber.Ctx) error {
		username := c.Params("username")
		data, err := unsplash.GetUserStats(username)
		if err != nil {
			return err
		}

		return c.JSON(data)
	})
}

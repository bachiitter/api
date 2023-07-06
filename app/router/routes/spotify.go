package routes

import (
	"time"

	"github.com/bachiitter/api/pkg/spotify"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
)

func SpotifyRouter(app fiber.Router) {
	router := app.Group("/spotify")

	router.Get("/player", func(c *fiber.Ctx) error {

		data, err := spotify.GetPlayer()
		if err != nil {
			return c.JSON(fiber.Map{
				"is_playing": false,
			})
		}

		return c.JSON(data)
	})

	router.Get("/tracks/now-playing", func(c *fiber.Ctx) error {

		data, err := spotify.GetNowPlaying()

		if err != nil {
			return c.JSON(fiber.Map{
				"is_playing": false,
			})
		}

		return c.JSON(data)
	})

	router.Get("/tracks/last-played", func(c *fiber.Ctx) error {

		limit := c.Query("limit")

		data, err := spotify.GetLastPlayed(limit)

		if err != nil {
			return c.JSON(err)
		}
		return c.JSON(data)
	})

	router.Use(cache.New(cache.Config{
		Expiration:   1 * time.Hour,
		CacheControl: true,
	}))

	router.Get("/tracks/top-tracks", func(c *fiber.Ctx) error {

		timeRange := c.Query("range")

		data, err := spotify.GetTopTracks(timeRange)
		if err != nil {
			return c.JSON(err)
		}

		return c.JSON(data)
	})
}

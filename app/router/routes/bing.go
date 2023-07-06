package routes

import (
	"github.com/bachiitter/api/pkg/bing"
	"github.com/gofiber/fiber/v2"
)

func BingRouter(app fiber.Router) {
	router := app.Group("/bing")

	router.Get("/time", func(c *fiber.Ctx) error {
		city := c.Query("city")

		data, err := bing.GetTime(city)
		if err != nil {
			return err
		}

		return c.JSON(data)
	})

	router.Get("/location", func(c *fiber.Ctx) error {
		location := c.Query("location")

		data, err := bing.GetLocation(location)
		if err != nil {
			return err
		}

		return c.JSON(data)
	})
}

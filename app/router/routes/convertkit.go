package routes

import (
	"github.com/bachiitter/api/pkg/convertkit"
	"github.com/gofiber/fiber/v2"
)

func ConverKitRouter(app fiber.Router) {
	router := app.Group("/convertkit")

	router.Post("/add-subscriber", func(c *fiber.Ctx) error {
		email := c.Query("email")

		data, err := convertkit.AddSubscriber(email)
		if err != nil {
			return err
		}

		return c.JSON(data)
	})
	router.Get("/total-subs", func(c *fiber.Ctx) error {
		data, err := convertkit.GetTotalSubs()
		if err != nil {
			return err
		}

		return c.JSON(data)
	})
}

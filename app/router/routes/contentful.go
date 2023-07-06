package routes

import (
	"github.com/bachiitter/api/pkg/contentful"
	"github.com/gofiber/fiber/v2"
)

func ContentfulRouter(app fiber.Router) {
	router := app.Group("/contentful")

	router.Get("/posts", func(c *fiber.Ctx) error {
		data, err := contentful.GetAllPosts()
		if err != nil {
			return c.JSON(err)

		}

		return c.JSON(data)
	})

	router.Get("/posts/:slug", func(c *fiber.Ctx) error {
		postSlug := c.Params("slug")

		data, err := contentful.GetPost(postSlug)
		if err != nil {
			return err

		}

		return c.JSON(data)
	})

	router.Get("/projects", func(c *fiber.Ctx) error {
		data, err := contentful.GetAllProjects()
		if err != nil {
			return err

		}

		return c.JSON(data)
	})
}

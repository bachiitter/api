package routes

import (
	"github.com/bachiitter/api/pkg/github"
	"github.com/gofiber/fiber/v2"
)

func GithubRouter(app fiber.Router) {
	router := app.Group("/github")

	router.Get("/repos/commits/:owner/:repo/:page", func(c *fiber.Ctx) error {

		owner := c.Params("owner")
		repo := c.Params("repo")
		page := c.Params("page")

		data, err := github.GetRepoCommits(owner, repo, page)
		if err != nil {
			return err
		}

		return c.JSON(data)
	})

	router.Get("/repos/:owner/:repo", func(c *fiber.Ctx) error {

		owner := c.Params("owner")
		repo := c.Params("repo")

		data, err := github.GetRepo(owner, repo)
		if err != nil {
			return err
		}

		return c.JSON(data)
	})

	router.Get("/users/:username", func(c *fiber.Ctx) error {

		user := c.Params("username")

		data, err := github.GetUser(user)
		if err != nil {
			return err
		}

		return c.JSON(data)
	})
}

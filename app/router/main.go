package router

import (
	"github.com/bachiitter/api/pkg"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	api.Get("/health", pkg.HandleHealthCheck)
}

package router

import (
	"github.com/bachiitter/api/app/router/routes"
	"github.com/bachiitter/api/pkg"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	api.Get("/health", pkg.HandleHealthCheck)

	routes.SpotifyRouter(api)
	routes.BingRouter(api)
	routes.ContentfulRouter(api)
	routes.ConverKitRouter(api)
	routes.GithubRouter(api)
	routes.OpenWeatherRouter(api)
	routes.UnsplashRouter(api)
}

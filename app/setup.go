package app

import (
	"os"
	"strconv"
	"time"

	"github.com/bachiitter/api/app/router"
	"github.com/bachiitter/api/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func New() error {

	// load env
	err := utils.LoadENV()
	if err != nil {
		return err
	}

	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())

	app.Use(compress.New(compress.Config{
		Next: func(c *fiber.Ctx) bool {
			a, _ := strconv.ParseBool(c.Get("x-no-compression"))
			return a
		},
		Level: compress.LevelBestSpeed, // 1
	}))

	var origins string

	goEnv := os.Getenv("GO_ENV")

	if goEnv == "" || goEnv == "development" {
		origins = "*"
	} else {

		origins = "https://bachitter.dev, https://portfolio-e3z.pages.dev, https://minimal.portfolio-e3z.pages.dev, https://v3.portfolio-e3z.pages.dev, http://localhost:3000"
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins: origins,
		AllowHeaders: "Content-Type",
		AllowMethods: "GET, POST",
		MaxAge:       600,
	}))

	app.Use(logger.New(logger.Config{
		Format: "${method} ${path} ${status}  ${latency}\n",
	}))

	app.Use(limiter.New(limiter.Config{
		Max:               200,
		Expiration:        5 * time.Minute,
		LimiterMiddleware: limiter.SlidingWindow{},
	}))

	// setup routes
	router.SetupRoutes(app)

	// Error Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Bruh! (>__<)")
	})

	// get the port and start
	port := os.Getenv("PORT")

	if port == "" {
		port = ":3000"
	} else {
		port = ":" + port
	}

	app.Listen(port)

	return nil
}

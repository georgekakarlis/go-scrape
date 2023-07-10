package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"goscrape.com/api/handlers"
	middleware "goscrape.com/middleware"
)

func SetupRoutes(app *fiber.App) {

	// Limit each IP to 10 requests per minute
    app.Use(limiter.New(limiter.Config{
        Max:        5,
        Expiration: 1 * time.Minute,
    }))

	app.Use(middleware.LoggerMiddleware)

	app.Get("/", func (c *fiber.Ctx) error  {
		return c.Render("index", fiber.Map{
			"Title": "hello world",
		})
	})

	//healthcheck
	app.Get("/api/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Never let your fears overcome your dreams",
		})
	})

	// POST to handle the form
	app.Post("/api/v1/process", handlers.ProcessForm)

	// GET to handle the download of the ready made file
	app.Get("/api/v1/download", handlers.DownloadFile)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})


}
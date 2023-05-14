package routes

import (
	"github.com/gofiber/fiber/v2"
	"goscrape.com/api/handlers"
)

func SetupRoutes(app *fiber.App) {

	//main route
	app.Get("/", handlers.IndexController)

	//healthcheck
	app.Get("/api/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Since you are curious, dm us k? or not :)",
		})
	})

	// POST to handle the form
	app.Post("/process", handlers.ProcessForm)

	// GET to handle the download of the ready made file
	app.Get("/download", handlers.DownloadCsvFile)
	app.Get("/download:filepath", handlers.DownloadCsvFile)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	// setup the user group
	//add the controllers		==TODO==
	/* user := app.Group("/user")
	user.Get("/")
	user.Post("/")
	user.Put("/:id")
	user.Get("/:id")
	user.Delete("/:id") */

	
}

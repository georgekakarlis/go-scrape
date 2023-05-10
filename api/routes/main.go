package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/goscrape/api/handlers"
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

	app.Post("/process", handlers.ProcessForm)


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

	//setup the scrape route group
	//controllers missing
	/* scrapeit := app.Group("/scrapeit")
	scrapeit.Get("/")
	scrapeit.Post("/")
	scrapeit.Put("/")
	scrapeit.Patch("/")
	scrapeit.Delete("/") */
}

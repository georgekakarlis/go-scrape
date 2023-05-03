package routes

import (
	"github.com/gofiber/fiber/v2"
)




func SetupRoutes(app *fiber.App) {
	 //healthcheck
	 app.Get("/api/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Since you are curious, dm us k? or not :)",
		})
	})

	// setup the user group
	//add the controllers		==TODO==
	user := app.Group("/user")
	user.Get("/")
	user.Post("/")
	user.Put("/:id")
	user.Get("/:id")
	user.Delete("/:id")

	//setup the scrape route group
	//controllers missing
	scrapeit := app.Group("/scrapeit")
	scrapeit.Get("/")
	scrapeit.Post("/")
	scrapeit.Put("/")
	scrapeit.Patch("/")
	scrapeit.Delete("/")
}
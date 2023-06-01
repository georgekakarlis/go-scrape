package routes

import (
	"github.com/gofiber/fiber/v2"
	"goscrape.com/api/handlers"
	middleware "goscrape.com/middlewares"
)

func SetupRoutes(app *fiber.App) {

	/* auth := app.Group("/auth")

    // User routes
    auth.Post("/register", handlers.Register)
    auth.Post("/login", handlers.Login)
    auth.Post("/logout", handlers.Logout)
    auth.Post("/refresh", handlers.RefreshToken) */

	//healthcheck
	
	app.Get("/api/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Since you are curious, dm us k? or not :)",
		})
	})

	// POST to handle the form
	app.Post("/api/process", handlers.ProcessForm)

	// GET to handle the download of the ready made file
	app.Get("/api/download", handlers.DownloadCsvFile)


	// Auth
	auth := app.Group("/auth")
	auth.Post("/login", handlers.Login)

	// User
	user := app.Group("/user")
	user.Get("/:id", handlers.GetUser)
	user.Post("/", handlers.CreateUser)
	user.Patch("/:id", middleware.Protected(), handlers.UpdateUser)
	user.Delete("/:id", middleware.Protected(), handlers.DeleteUser)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404) // => 404 "Not Found"
	})
	

}

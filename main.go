package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"goscrape.com/api/routes"
	"goscrape.com/initializers"
	"goscrape.com/middlewares"

	"github.com/gofiber/template/html/v2"
)

func init() {
	initializers.ConnectDB()
}

func main() {

	//load templates
	engine := html.New("./views", ".html")
	// Reload the templates on each render, good for development
	engine.Reload(true)

	// Create new Fiber instance
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	//set middlewares
	middlewares.SetMiddlewares(app)
	// setup routes
	routes.SetupRoutes(app)

	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	//serve
	log.Printf("Starting up on http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}

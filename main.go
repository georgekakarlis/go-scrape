package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"goscrape.com/api/routes"
	"goscrape.com/initializers"
	"goscrape.com/middlewares"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func init() {
	initializers.ConnectDB()

}

func main() {

	// Create new Fiber instance
	app := fiber.New()


	//set middlewares
	middlewares.SetMiddlewares(app)
	// setup routes
	routes.SetupRoutes(app)

	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}


	//serve
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
       log.Fatal(err)
    }
    defer file.Close()
    log.SetOutput(file)
	log.Printf("Starting up on http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}

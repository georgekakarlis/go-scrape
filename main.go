package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/helmet/v2"
	"github.com/goscrape/api/routes"
	"github.com/goscrape/initializers"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func init() {
	initializers.ConnectDB()
}

func main() {
 
  // Create new Fiber instance
	app := fiber.New()
 	
	// attach middlewares
	app.Use(recover.New())	//Recover middleware for Fiber that recovers from panics anywhere in the stack chain and handles the control to the centralized ErrorHandler.
	app.Use(logger.New(logger.Config{		////Logger middleware for Fiber that logs HTTP request/response details.
		Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
	}))		
	app.Use(helmet.New())		//helmet middleware :)

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

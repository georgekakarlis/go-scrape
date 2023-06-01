package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"goscrape.com/api/routes"
	"goscrape.com/database"
	"goscrape.com/logger"
	middleware "goscrape.com/middlewares"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)



func main() {

	// Initialize the Log DB
	loggerErr := logger.InitializeDB()
	if loggerErr != nil {
		log.Fatal("💾  Failed to connect to the Log Database! \n", loggerErr.Error())
		os.Exit(1)
	}

	// Create new Fiber instance
	app := fiber.New()

	// db connec t
	database.ConnectDB()

	//set middlewares
	middleware.SetMiddlewares(app)
	
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
	log.Printf("🤖  Starting up on http://localhost:%s", port)
	fmt.Printf(" 🤖Starting up on http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}

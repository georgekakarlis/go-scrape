package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"goscrape.com/api/routes"
	"goscrape.com/database"
	"goscrape.com/initializers"
	"goscrape.com/logger"
	"goscrape.com/middleware"
)

var (
    WarningLogger *log.Logger
    InfoLogger    *log.Logger
    ErrorLogger   *log.Logger
)

func init () {
	initializers.LoadEnvVariables()
	// db connec t
	database.ConnectDB()
}


func main() {

	// Initialize the Log DB
	loggerErr := logger.InitializeDB()
	if loggerErr != nil {
		log.Fatal("💾  Failed to connect to the Log Database! \n", loggerErr.Error())
		os.Exit(1)
	}

	// Create new Fiber instance
	router := gin.Default()


	

	//set middlewares
	middleware.SetMiddlewares(router)
	
	// setup routes
	routes.SetupRoutes(router)

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
	log.Fatal(router.Run(":" + port))
}

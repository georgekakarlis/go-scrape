package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/goscrape/initializers"
)

func init() {
	initializers.ConnectDB()
}

func main() {
  
	//db connect	 
	

  // Create new Fiber instance
	app := fiber.New()
 	
	//main router
  

  //middlewares
  


  // base route
 
  //healthcheck
  app.Get("/api/healthchecker", func(c *fiber.Ctx) error {
	return c.Status(200).JSON(fiber.Map{
		"status":  "success",
		"message": "Since you are curious, dm us k? or not :)",
	})
})

  
  //routes declaration
  

  	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
    port = fromEnv
  }
	//serve
	log.Printf("Starting up on http://localhost:%s", port)
	log.Fatal(app.Listen(":" + port))
}

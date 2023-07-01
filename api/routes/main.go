package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"goscrape.com/api/handlers"
)

func SetupRoutes(router *gin.Engine) {
	
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowCredentials = true

	// Apply CORS middleware
	router.Use(cors.New(config))
	//router.GET("/" , middleware.RequireAuth)

	/* router.POST("/api/v1/signup", handlers.Signup)
	router.POST("/api/v1/login", handlers.Login)
	router.GET("/api/v1/validate", middleware.RequireAuth, handlers.Validate)
	router.POST("/api/v1/logout", handlers.Logout) */
	

	// POST to handle the form
	router.POST("/api/v1/process",  handlers.ProcessForm)

	// GET to handle the download of the ready made file
	router.GET("/api/v1/download",  handlers.DownloadCsvFile)


	// 404 Handler
	router.NoRoute(func(c *gin.Context) {
		c.Status(404) // => 404 "Not Found"
	})
	

}



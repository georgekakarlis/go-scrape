package routes

import (
	"github.com/gin-gonic/gin"
	"goscrape.com/api/handlers"
	"goscrape.com/middleware"
)

func SetupRoutes(router *gin.Engine) {

	router.GET("/")

	router.POST("/api/v1/signup", handlers.Signup)
	router.POST("/api/v1/login", handlers.Login)
	router.GET("/api/v1/validate", middleware.RequireAuth, handlers.Validate)

	// POST to handle the form
	router.POST("/api/v1/process", handlers.ProcessForm)

	// GET to handle the download of the ready made file
	router.GET("/api/v1/download", handlers.DownloadCsvFile)


	// 404 Handler
	router.NoRoute(func(c *gin.Context) {
		c.Status(404) // => 404 "Not Found"
	})
	

}



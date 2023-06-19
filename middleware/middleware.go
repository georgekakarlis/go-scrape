package middleware

import (
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



func SetMiddlewares(router *gin.Engine) {
	// Attach middlewares

	// Add a logger middleware with custom formatter
	router.Use(func(c *gin.Context) {
		start := time.Now()

		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		log.Printf("[ %s ] %s - %s %s %d %s\n",
			c.ClientIP(),
			end.Format(time.RFC3339),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
		)
	})
	//cors
			config := cors.DefaultConfig() // defaultconfig doesnot allow all origins, we still have to define which
  			config.AllowOrigins = []string{"http://localhost:3000"}
  	router.Use(cors.New(config))
}



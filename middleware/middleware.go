package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// LoggerMiddleware is a middleware that logs request details
func LoggerMiddleware(ctx *fiber.Ctx) error {
	start := time.Now()

	err := ctx.Next()

	end := time.Now()
	latency := end.Sub(start)

	log.Printf("[ %s ] %s - %s %s %d %s\n",
		ctx.IP(),
		end.Format(time.RFC3339),
		ctx.Method(),
		ctx.Path(),
		ctx.Response().StatusCode(),
		latency,
	)

	return err
}

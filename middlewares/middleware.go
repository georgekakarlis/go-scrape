package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)



func SetMiddlewares(app *fiber.App) {
		// attach middlewares
		app.Use(recover.New())            //Recover middleware for Fiber that recovers from panics anywhere in the stack chain and handles the control to the centralized ErrorHandler.
		app.Use(logger.New(logger.Config{ ////Logger middleware for Fiber that logs HTTP request/response details.
			Format: "[${ip}]:${port} ${status} - ${method} ${path} ${latency}\n",
		}))
		app.Use(helmet.New()) //helmet middleware :)
		app.Use(cors.New())
}
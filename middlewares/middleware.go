package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"goscrape.com/config"
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

// Protected protect routes
//This middleware function is responsible for protecting routes by enforcing JWT authentication. 
//It returns a fiber.Handler, which can be used as middleware in Fiber to protect routes that require authentication. 
//Inside this function, jwtware.New() is used to create a new JWT middleware with the provided configuration. 
//The configuration includes the SigningKey (which is the secret key used to sign and verify the JWT) retrieved from the configuration file using config.Config("SECRET"). 
//This middleware will verify the JWT token provided in the request headers and ensure its validity. If the token is valid, the request will proceed to the next handler in the chain. If the token is missing, malformed, or expired, the jwtError function will be called to handle the error.
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(config.Config("SECRET")),
		ErrorHandler: jwtError,
	})
}

//This function is the error handler for the JWT middleware. It takes a Fiber context (c) and an error (err) as parameters. 
//Based on the error received, it generates an appropriate error response. If the error message is "Missing or malformed JWT", it means that the JWT token is either missing or not properly formatted. 
//In this case, the function returns a JSON response with a status code of fiber.StatusBadRequest (400 Bad Request) and an error message indicating the issue. 
//For any other error, it assumes that the JWT is either invalid or expired and returns a JSON response with a status code of fiber.StatusUnauthorized (401 Unauthorized) and an error message indicating the issue.
func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
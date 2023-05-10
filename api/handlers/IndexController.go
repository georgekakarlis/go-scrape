package handlers

import (
	"github.com/gofiber/fiber/v2"
)


func IndexController (c *fiber.Ctx) error {
	  // Render index
	  return c.Render("index", fiber.Map{
		"Title": "Hello, <b>World</b>!",
	})
}
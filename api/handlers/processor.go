package handlers

import (
	"github.com/gofiber/fiber/v2"
	"goscrape.com/helpers"
	"goscrape.com/scrape"
)

// ProcessForm handles the form submission
func ProcessForm(c *fiber.Ctx) error {
	// Retrieve the form data
	form := new(struct {
		URL          string `form:"url"`
		GenerateFILE string `form:"generateFILE"`
	})

	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid form data")
	}

	// Validate the URL entered by the user
	// further regex validation ==TODO==
	if form.URL == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Please enter a URL")
	}

	// Scrape the URL
	scrapedData := scrape.ScrapeURL(form.URL)

	// Check the output format requested by the user
	switch form.GenerateFILE {
	case "generateCSV":
		// Generate CSV output
		fileName, err := helpers.MakeCSV(scrapedData)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Failed to generate CSV")
		}
		return c.JSON(fiber.Map{"filePath": fileName})

	case "generateXLSX":
		// generate Excel output

	case "generatePDF":
		// Generate PDF output and
		// replace the return statement with the appropriate code

	case "generateJSON":
		// Default to JSON output
		return c.JSON(scrapedData)
	}

	// Handle an invalid or unsupported format
	return c.Status(fiber.StatusBadRequest).SendString("Invalid format selected")

}

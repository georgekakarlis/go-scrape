package handlers

import (
	"net/url"
	"regexp"

	"github.com/gofiber/fiber/v2"
	"goscrape.com/helpers"
	"goscrape.com/logger"
	"goscrape.com/scrape"
)

// ProcessForm handles the form submission
func ProcessForm(c *fiber.Ctx) error {
	// Retrieve the form data
	form := new(struct {
		URL          string `json:"url"`
		GenerateFILE string `json:"generateFILE"`
		//SpecifiedItem string `json:"specifiedItem"`
	})

	if err := c.BodyParser(form); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid form data")
	}

	// Validate the URL entered by the user
	if form.URL == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Please enter a valid URL")
	}		

	// Check if the URL is valid regex
	if !ValidateURL(form.URL) {
		return c.Status(fiber.StatusBadRequest).SendString("Please enter a valid URL")
	}

	// Scrape the URL
	scrapedData := scrape.ScrapeURL(form.URL)

	// Log the request details and save to the sqlite
	logger.LogRequest(c.IP(), form.URL)



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


	case "generateJSON":
		// Default to JSON output
		return c.JSON(scrapedData)
	}

	// Handle an invalid or unsupported format
	return c.Status(fiber.StatusBadRequest).SendString("Invalid format selected")
}
// ^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|\/|\/\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$
// ^(https?|ftp|file):\/\/[-A-Za-z0-9+&@#\/%?=~_|!:,.;]*[-A-Za-z0-9+&@#\/%=~_|]
func ValidateURL(inputURL string) bool {
	// Check if the URL is empty
	if inputURL == "" {
		return false
	}

	// Check if the URL is a valid format
	_, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return false
	}

	// Regex pattern for validating the URL
	regexPattern := `^(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|\/|\/\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?$`

	// Compile the regex pattern
	regex, err := regexp.Compile(regexPattern)
	if err != nil {
		return false
	}

	// Match the regex pattern against the input URL
	match := regex.MatchString(inputURL)

	return match
}



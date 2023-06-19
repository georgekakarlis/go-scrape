package handlers

import (
	"net/http"
	"net/url"
	"regexp"

	"github.com/gin-gonic/gin"
	"goscrape.com/helpers"
	"goscrape.com/logger"
	"goscrape.com/scrape"
)

// ProcessForm handles the form submission
func ProcessForm(router *gin.Context) {
	// Retrieve the form data
	form := struct {
		URL          string `json:"url"`
		GenerateFILE string `json:"generateFILE"`
	}{}
		//shouldbindjson is the gin's bodyparser somehow
	if err := router.ShouldBindJSON(&form); err != nil {
		router.String(http.StatusBadRequest, "Invalid form data")
		return
	}

	// Validate the URL entered by the user
	if form.URL == "" {
		router.String(http.StatusBadRequest, "Please enter a valid URL")
		return
	}

	// Check if the URL is valid regex
	if !ValidateURL(form.URL) {
		router.String(http.StatusBadRequest, "Please enter a valid URL")
		return
	}

	// Scrape the URL
	scrapedData := scrape.ScrapeURL(form.URL)

	// Log the request details and save to the database
	logger.LogRequest(router.ClientIP(), form.URL)

	// Check the output format requested by the user
	switch form.GenerateFILE {
	case "generateCSV":
		// Generate CSV output
		fileName, err := helpers.MakeCSV(scrapedData)
		if err != nil {
			router.String(http.StatusInternalServerError, "Failed to generate CSV")
			return
		}
		router.JSON(http.StatusOK, gin.H{"filePath": fileName})

	case "generateXLSX":
		// generate Excel output

	case "generatePDF":
		// Generate PDF output and
		// replace the return statement with the appropriate code

	case "generateJSON":
		// Default to JSON output
		router.JSON(http.StatusOK, scrapedData)

	default:
		// Handle an invalid or unsupported format
		router.String(http.StatusBadRequest, "Invalid format selected")
	}
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



package handlers

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
)




func DownloadCsvFile(c *fiber.Ctx) error {

	filePath := c.Params("filepath")
    
	// Generate a unique filename for the CSV file
	filename := fmt.Sprintf("%s.csv", time.Now().Format("2006-01-02_15-04-05"))

	// Set the response headers for file download
	c.Set(fiber.HeaderContentType, "text/csv")
	c.Set(fiber.HeaderContentDisposition, fmt.Sprintf(`attachment; filename="%s"`, filename))

	// Serve the CSV file from the temporary directory
	return c.SendFile(filepath.Join("dir", filename, filePath))

}


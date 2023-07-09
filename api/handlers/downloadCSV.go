package handlers

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func DownloadCsvFile(c *fiber.Ctx) error {
	filePath := c.Query("filePath")

	// Check if the filePath is empty
	if filePath == "" {
		return c.Status(fiber.StatusBadRequest).SendString("invalid file path")
	}

	// Open the file to get its size
	file, err := os.Open(filePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to open file")
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("failed to get file info")
	}
	fileSize := fileInfo.Size()

	// Set the response headers
	c.Set(fiber.HeaderContentType, "text/csv")
	c.Set(fiber.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", filepath.Base(filePath)))
	c.Set(fiber.HeaderContentLength, strconv.FormatInt(fileSize, 10))

	// Send the file as the response body
	if err := c.SendFile(filePath); err != nil {
		return err // Return the error directly
		}

	// Delete the file
	if err := os.Remove(filePath); err != nil {
		fmt.Printf("Failed to delete file: %v\n", err)
	}

		return nil

}
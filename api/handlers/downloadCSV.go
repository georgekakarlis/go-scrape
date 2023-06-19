package handlers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DownloadCsvFile(router *gin.Context) {
	filePath := router.Query("filePath")

	// Check if the filePath is empty
	if filePath == "" {
		router.String(http.StatusBadRequest, "invalid file path")
		return
	}

	// Open the file to get its size
	file, err := os.Open(filePath)
	if err != nil {
		router.String(http.StatusInternalServerError, "failed to open file")
		return
	}
	defer file.Close()

	// Get the file size
	fileInfo, err := file.Stat()
	if err != nil {
		router.String(http.StatusInternalServerError, "failed to get file info")
		return
	}
	fileSize := fileInfo.Size()

	// Set the response headers
	router.Header("Content-Type", "text/csv")
	router.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filepath.Base(filePath)))
	router.Header("Content-Length", strconv.FormatInt(fileSize, 10))

	// Send the file as the response body
	router.File(filePath)

	// Delete the file
	if err := os.Remove(filePath); err != nil {
		fmt.Printf("Failed to delete file: %v\n", err)
	}
}
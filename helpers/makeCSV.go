package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// MakeCSV function will help us create a function that is gonna create a csv file from the scraped data
//the ioutil.TempFile function is used to create a temporary file with the prefix "scraped_data" and the .csv extension.
//The TempFile function automatically generates a unique name for the temporary file.


func MakeCSV(scrapedData []string, directory string) (string, error) {
	// Ensure that the directory exists
	if _, err := os.Stat("directory"); os.IsNotExist(err) {
		// Directory does not exist, create it
		err := os.Mkdir("directory", 0755) // Specify the desired directory permissions (e.g., 0755)
		if err != nil {
			fmt.Println("Failed to create directory:", err)
			return "", err
		}
	}

	filePath := filepath.Join(directory, "scraped_data.csv")

	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return "", err
	}

	defer func() {
		_ = file.Close()
		_ = os.Remove(file.Name()) // Remove the temporary file if not already deleted
	}()

	// Write scrapedData to the file
	content := []byte(strings.Join(scrapedData, "\n"))
	_, err = file.Write(content)
	if err != nil {
		fmt.Println("Failed to write to file:", err)
		return "", err
	}

	return file.Name(), nil
}

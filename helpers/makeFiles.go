package helpers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

//The GenerateCSV, GenerateXLSX, and GenerateJSON functions handle the generation of respective file formats, while the ensureDirectory and moveFile functions assist with directory creation and file moving operations, respectively.
//By organizing the code into helper functions, the route handler becomes cleaner and more focused on handling the request and response.

// GenerateCSV generates a CSV file from the scraped data and returns the file path
func GenerateCSV(data []string) (string, error) {
	fileName, err := MakeCSV(data)
	if err != nil {
		return "", err
	}

	csvDir := "./downloads/CSV"
	if err := ensureDirectory(csvDir); err != nil {
		return "", err
	}

	destPath := filepath.Join(csvDir, "scraped-data.csv")
	if err := moveFile(fileName, destPath); err != nil {
		return "", err
	}

	return destPath, nil
}

// GenerateXLSX generates an XLSX file from the scraped data and returns the file path
func GenerateXLSX() {
	/* fileName := fmt.Sprintf("%s.xlsx", time.Now().Format("2006-01-02_15-04-05"))
	xlsxDir := "./downloads/XLSX"
	if err := ensureDirectory(xlsxDir); err != nil {
		return "", err
	}

	filePath := filepath.Join(xlsxDir, fileName)
	f := excelize.NewFile()
	sheetName := "New Scraped Data"
	sheetIndex := f.NewSheet(sheetName)
	for i, link := range data {
		cellName := fmt.Sprintf("A%d", i+1)
		f.SetCellValue(sheetName, cellName, link)
	}
	f.SetActiveSheet(sheetIndex)

	if err := f.SaveAs(filePath); err != nil {
		return "", err
	}

	return filePath, nil */
}

// GenerateJSON generates a JSON file from the scraped data and returns the file path
func GenerateJSON(data []string) (string, error) {
	fileName := fmt.Sprintf("%s.json", time.Now().Format("2006-01-02_15-04-05"))
	jsonDir := "./downloads/JSON"
	if err := ensureDirectory(jsonDir); err != nil {
		return "", err
	}

	filePath := filepath.Join(jsonDir, fileName)

	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	if err := ioutil.WriteFile(filePath, jsonData, 0644); err != nil {
		return "", err
	}

	return filePath, nil
}


// moveFile moves a file from sourcePath to destPath
func moveFile(sourcePath, destPath string) error {
	return os.Rename(sourcePath, destPath)
}
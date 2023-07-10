package helpers

import (
	"os"
	"time"

	"github.com/tealeg/xlsx"
)

func MakeXLSX(data []string) (string, error) {
    var file *xlsx.File
    var sheet *xlsx.Sheet
    var row *xlsx.Row
    var cell *xlsx.Cell
    var err error

    file = xlsx.NewFile()
    sheet, err = file.AddSheet("Sheet1")
    if err != nil {
        return "", err
    }

    for _, url := range data {
        row = sheet.AddRow()
        cell = row.AddCell()
        cell.Value = url
    }

    // Ensure the directory exists
    if _, err := os.Stat("./downloads/XLSX"); os.IsNotExist(err) {
        // Directory does not exist, create it
        os.Mkdir("./downloads/XLSX", 0755)
    }

    filename := "./downloads/XLSX/scraped_" + time.Now().Format("20060102150405") + ".xlsx"
    err = file.Save(filename)
    if err != nil {
        return "", err
    }

    return filename, nil
}
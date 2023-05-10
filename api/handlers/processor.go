package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/goscrape/scrape"

	"github.com/xuri/excelize/v2"
)

func ProcessForm(c *fiber.Ctx) error {

	url := c.FormValue("url")
	format := c.Query("format", "json")

    scrapedData := scrape.ScrapeURL(url)

    // Check the output format requested by the user
    switch format {
    case "csv":
        // Generate CSV output
        return c.SendString(strings.Join(scrapedData, ","))
	case "xlsx":
		// Generate XLSX output
		f := excelize.NewFile()
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Println(err)
			}
		}()
		sheetName := "Sheet1"
		sheetIndex, err := f.NewSheet(sheetName)
		if err != nil {
			fmt.Println(err)
			return err
		}
		for i, link := range scrapedData {
			cellName := fmt.Sprintf("A%d", i+1)
			f.SetCellValue(sheetName, cellName, link)
		}
		f.SetActiveSheet(sheetIndex)
		filename := fmt.Sprintf("%s.xlsx", time.Now().Format("2006-01-02_15-04-05"))
		return c.SendFile(f, filename, fmt.Sprintf("attachment; filename=%s", filename))
	
	
    default:
        // Default to JSON output
        return c.JSON(scrapedData)
    }
}

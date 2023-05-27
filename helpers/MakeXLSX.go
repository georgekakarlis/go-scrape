package helpers

import (
	

	"github.com/xuri/excelize/v2"
)

//this function creates the excel file.
//poor docs so I have to think of it sth useful



func MakeXLSX(scrapedData [][]string) error {
    file := excelize.NewFile()
   

  

   

    // Save the Excel file
    if err := file.SaveAs("output.xlsx"); err != nil { // Provide the desired file name and path
        return err
    }

    return nil
}
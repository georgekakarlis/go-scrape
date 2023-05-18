package helpers

import (
	"encoding/csv"
	"log"
	"os"
)

// MakeCSV function will help us create a function that is gonna create a csv file from the scraped data
//the ioutil.TempFile function is used to create a temporary file with the prefix "scraped_data" and the .csv extension.
//The TempFile function automatically generates a unique name for the temporary file.

/* // Ensure that the directory exists
if _, err := os.Stat("directory"); os.IsNotExist(err) {
	// Directory does not exist, create it
	err := os.Mkdir("directory", 0755) // Specify the desired directory permissions (e.g., 0755)
	if err != nil {
		fmt.Println("Failed to create directory:", err)
		return "", err
	}
} */
	func MakeCSV(scrapedData []string, ) (string, error) {

	// file name and location of our csv file 
	var fileName string = "./downloads/CSV/scraped-data.csv"
	
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal("Panic: Not able to create file. \n", fileName, err)
	}
	
	defer file.Close()

	// writer will write the context of the file
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// first two heading of the CSV file
	writer.Write(scrapedData)
	

	return fileName, err
		
	
}
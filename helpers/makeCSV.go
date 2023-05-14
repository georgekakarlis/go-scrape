package helpers

import (
	"encoding/csv"
	"log"
	"os"
)

//makeCSV function will help us create a function that is gonna create a csv file from the scraped data


func makeCSV() {
	filename := "your-data.csv"
    file, err := os.Create(filename)
    if err != nil {
     log.Fatalf("could not create the file, err :%q",err)
     return
    }
    defer file.Close()

	//The next thing we do with a writer once we are done writing the file, 
	//we throw everything from the buffer into the writer, which can later be passed onto the file. For that, we will use Flush.
	//This process has to be performed afterward and not right away. So, we add the keyword defer.
	writer := csv.NewWriter(file)
	defer writer.Flush()
}
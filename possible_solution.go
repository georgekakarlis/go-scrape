The error messages indicate that you are passing the wrong types of arguments to the SendFile method. According to the Fiber documentation, the SendFile method expects a string argument for the file path and an optional bool argument for whether to enable inline mode. However, you are passing a *excelize.File pointer as the first argument and a string as the second argument. This causes a type mismatch and a compilation error.

To fix this, you need to save the xlsx file to a temporary location and pass the file path as a string to the SendFile method. You can also pass true or false as the third argument to enable or disable inline mode. For example:

// Generate XLSX output
f := excelize.NewFile()
defer func() {
if err := f.Close(); err != nil {
fmt.Println(err)
}
}()
sheetName := "New Scraped Data"
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
// Save the xlsx file to a temporary location
tmpfile, err := ioutil.TempFile("", filename)
if err != nil {
fmt.Println(err)
return err
}
defer os.Remove(tmpfile.Name()) // clean up
if err := f.SaveAs(tmpfile.Name()); err != nil {
fmt.Println(err)
return err
}
// Send the xlsx file as an HTTP response
return c.SendFile(tmpfile.Name(), true) // enable inline mode

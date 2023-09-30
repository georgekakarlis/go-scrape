package scrape

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/extensions"
)


func ScrapeURL(url string, ) []string {

	//MAKE THIS CODE WITH A SWITCH CASE TO SCRAPE BASED ON AHREF, DIV, P , Hs AND SO ON. 

	//colly
	c := colly.NewCollector()


    extensions.RandomUserAgent(c)
  

	// Create a slice to hold the scraped data because slices are built on top of arrays and we dont know what or how much we expect to get back from scraping
	var scrapedData []string

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		scrapedData = append(scrapedData, e.Attr("href"))
	})

	// set a valid User-Agent header
	//c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"

	// timeout on request
	c.SetRequestTimeout(120 * time.Second)

	// where are u going colly?
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting:", r.URL)
	})

	// what did u get back?
	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Got a response from", r.Request.URL)
	})

	//oh u got an error
	c.OnError(func(r *colly.Response, e error) {
		fmt.Println("Got this error:", e)
	})	
	
	c.Visit(url)

	return scrapedData

}



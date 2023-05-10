package scrape

import (
	"github.com/gocolly/colly/v2"
)

type pageInfo struct {
	StatusCode int
	Links	map[string]int
}

func ScrapeURL(url string) []string {

	/* URL := r.URL.Query().Get("url")
	if URL == "" {
		log.Println("missing URL argument")
		return
	}
	log.Println("visiting", URL) */


	//colly
	c := colly.NewCollector()

	// Create a slice to hold the scraped data
    var scrapedData []string

	

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        scrapedData = append(scrapedData, e.Attr("href"))
    })

	 // set a valid User-Agent header
	 c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36"


	
	

	c.Visit(url)

	return scrapedData

}

package scrape

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)


func scrape() {
	//colly
	c := colly.NewCollector(colly.AllowedDomains("www.mydomain.com"))

	c.OnHTML("a[href]", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})

	c.Visit("www.mydomain.com/somefunny/url")

}
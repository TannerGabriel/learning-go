package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("gabrieltanner.org"),
	)

	// Callback for when a scraped page contains an article container
	c.OnHTML(".article-container", func(e *colly.HTMLElement) {
		fmt.Println("Article heading: ", e.DOM.Find("h1").Text())
	})

	// Callback for links on scraped pages
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Extract the linked URLs from the anchor tag
		link := e.Attr("href")
		// Have your crawler visit the linked URL
		c.Visit(e.Request.AbsoluteURL(link))
	})

	// Callback getting called when the scraping process is finished
	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished scraping!")
	})

	// Set a rate limit for the crawler
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 1 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.Visit("https://gabrieltanner.org")
}

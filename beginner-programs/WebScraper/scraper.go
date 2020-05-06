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

    // Callback for when a scraped page contains an article heading
    c.OnHTML(".article-heading", func(e *colly.HTMLElement) {   
		fmt.Println("Article heading: ", e.Text)     
    })

    // Callback for links on scraped pages
    c.OnHTML("a[href]", func(e *colly.HTMLElement) {
        // Extract the linked URLs from the anchor tag
        link := e.Attr("href")
        // Have your crawler visit the linked URL
        c.Visit(e.Request.AbsoluteURL(link))
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
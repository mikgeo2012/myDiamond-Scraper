package main

import (

	"myDiamond-scraper/pkg/models"
	"fmt"
	"myDiamond-scraper/pkg/scraper"
)

func main() {
	o := models.GetDefault()
	if o.Validate() {
		fmt.Println("Valid")
	} else {
		fmt.Println("Not Valid")
	}

	w := scraper.WebScraper{}
	w.Initialize(o)
}

// Scrape each page of James Allen
// Parse
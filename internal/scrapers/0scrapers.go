package scrapers

import "github.com/gocolly/colly"

var scraper *colly.Collector

func Init() {
	scraper = colly.NewCollector(colly.Async(true))
}

func Get() *colly.Collector {
	return scraper
}

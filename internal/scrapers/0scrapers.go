package scrapers

import (
	"go-scrapers/internal/scrapers/amazon"
	"go-scrapers/internal/scrapers/worldometers"

	"github.com/gocolly/colly/v2"
)

var scraper *colly.Collector

func Init() {
	scraper = colly.NewCollector(
		colly.Async(true),
	)
}

func InitAmazon(query string) {
	amazon.Scrape(scraper.Clone(), query)
}

func InitWorldometer() {
	worldometers.Scrape(scraper.Clone())
}

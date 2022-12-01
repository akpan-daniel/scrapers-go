package scrapers

import (
	"go-scrapers/internal/scrapers/amazon"
	"go-scrapers/internal/scrapers/arizona-real-estate"
	"go-scrapers/internal/scrapers/worldometers"

	"github.com/gocolly/colly/v2"
)

var scraper *colly.Collector

func Init() {
	scraper = colly.NewCollector()
}

func InitAmazon(query string) {
	amazon.Scrape(scraper.Clone(), query)
}

func InitWorldometer() {
	worldometers.Scrape(scraper.Clone())
}

func InitArizona() {
	arizona.Scrape(scraper.Clone())
}

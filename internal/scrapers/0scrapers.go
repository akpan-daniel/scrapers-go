package scrapers

import (
	"go-scrapers/internal/scrapers/worldometers"

	"github.com/gocolly/colly/v2"
)

func Init(query string) {
	scraper := colly.NewCollector(
		colly.Async(true),
	)

	// amazon.Scrape(scraper.Clone(), query)
	worldometers.Scrape(scraper.Clone())
}

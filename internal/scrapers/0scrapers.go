package scrapers

import (
	"github.com/gocolly/colly"
	"go-scrapers/internal/scrapers/amazon"
)

func Init(query string) {
	scraper := colly.NewCollector(
		colly.Async(true),
	)

	amazon.Scrape(scraper.Clone(), query)
}

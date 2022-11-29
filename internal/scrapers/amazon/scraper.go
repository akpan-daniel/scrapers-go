package amazon

import (
	"encoding/csv"
	"fmt"
	"go-scrapers/internal/utils"
	"log"

	"github.com/gocolly/colly/v2"
)

var (
	ScraperDomain = "www.amazon.com"
	DirName       = "amazon/"
	failedUrls    = make(map[string]string)
)

func Scrape(scraper *colly.Collector, query string) {
	utils.MakeDirs(&DirName)

	file := utils.OpenFile(DirName + "result.csv")
	defer utils.CloseFile(file)

	defer utils.LogFailedUrls(DirName, failedUrls)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Create Headers
	_ = writer.Write(Product{}.GetHeaders())

	utils.RegisterHandlers(scraper, DirName, ScraperDomain)

	scraper.OnError(func(response *colly.Response, err error) {
		url := response.Request.URL.String()
		failedUrls[url] = err.Error()
	})

	scraperURL := fmt.Sprintf("https://%s/s?k=%s", ScraperDomain, query)

	detailScraper := scraper.Clone()
	utils.RegisterHandlers(detailScraper, DirName, ScraperDomain)

	scraper.OnHTML("div[data-component-type='s-search-result']", func(element *colly.HTMLElement) {
		productURL := element.ChildAttr("h2.a-size-mini > a", "href")
		next := element.Request.AbsoluteURL(productURL)

		if err := detailScraper.Visit(next); err != nil {
			failedUrls[next] = err.Error()
		}
	})

	scraper.OnHTML("a.s-pagination-next", func(element *colly.HTMLElement) {
		url := element.Request.AbsoluteURL(element.Attr("href"))
		if err := scraper.Visit(url); err != nil {
			failedUrls[url] = err.Error()
		}
	})

	detailScraper.OnHTML("#dp-container", func(element *colly.HTMLElement) {
		product := Product{
			Title:     element.ChildText("h1#title > span"),
			Price:     element.ChildText("#ppd .a-price[data-a-size='xl'] > .a-offscreen"),
			URL:       element.Request.URL.String(),
			Seller:    element.ChildText("#sellerProfileTriggerId"),
			Rating:    element.ChildText("##acrPopover .a-declarative .a-popover-trigger .a-icon-alt:nth-child(1)"),
			Questions: element.ChildText("#askATFLink"),
			Reviews:   element.ChildText("#acrCustomerReviewText"),
		}
		if err := writer.Write(product.ToSlice()); err != nil {
			log.Fatalln(err)
		}
	})

	if err := scraper.Visit(scraperURL); err != nil {
		log.Fatalf("Unable to start scraper:\n%v", err)
	}

	scraper.Wait()
	detailScraper.Wait()
}

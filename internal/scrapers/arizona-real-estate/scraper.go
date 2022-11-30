package arizona

import (
	"encoding/csv"
	"fmt"
	"go-scrapers/internal/utils"
	"log"
	"time"

	"github.com/gocolly/colly/v2"
)

var (
	ScraperDomain = "www.arizonarealestate.com"
	DirName       = "arizona-real-estates"
	failedUrls    = make(map[string]string)
)

func Scrape(scraper *colly.Collector) {
	utils.MakeDirs(&DirName)

	file := utils.OpenFile(DirName + "result.csv")
	defer utils.CloseFile(file)

	defer utils.LogFailedUrls(DirName, failedUrls)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	_ = writer.Write(Property{}.GetHeaders())

	utils.RegisterHandlers(scraper, DirName, ScraperDomain)

	scraper.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 4,
		Delay:       5 * time.Second,
	})

	count := 0

	scraper.OnXML("//div[@class='si-listings-column']", func(element *colly.XMLElement) {
		property := Property{
			Name:        element.ChildText(".//div[@class='si-listing__title-main']"),
			Description: element.ChildText(".//div[@class='si-listing__title-description']"),
			Agency:      element.ChildText(".//div[@class='si-listing__footer']/div"),
			Price:       element.ChildText(".//div[@class='si-listing__photo-price']/span"),
		}
		_ = writer.Write(property.ToSlice())
		count++
	})

	scraper.OnXML("//li[@class='next']", func(element *colly.XMLElement) {
		url := element.ChildAttr(".//a", "href")
		nextUrl := element.Request.AbsoluteURL(url)

		if err := scraper.Visit(nextUrl); err != nil {
			failedUrls[url] = err.Error()
			log.Println("Error visiting:", err)
		}
	})

	scraperURL := fmt.Sprintf("https://%s/maricopa/", ScraperDomain)

	if err := scraper.Visit(scraperURL); err != nil {
		log.Fatalln("Unable to start scraper:", err.Error())
	}

	scraper.Wait()
}

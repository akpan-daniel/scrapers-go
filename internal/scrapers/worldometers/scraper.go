package worldometers

import (
	"encoding/csv"
	"fmt"
	"go-scrapers/internal/utils"
	"log"

	"github.com/gocolly/colly/v2"
)

var (
	ScraperDomain = "www.worldometers.info"
	DirName       = "worldometer/"
	failedUrls    = make(map[string]string)
)

func Scrape(scraper *colly.Collector) {
	scraper.MaxDepth = 2

	utils.MakeDirs(&DirName)

	file := utils.OpenFile(DirName + "result.csv")
	defer utils.CloseFile(file)

	defer utils.LogFailedUrls(DirName, failedUrls)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	_ = writer.Write(Population{}.GetHeaders())

	utils.RegisterHandlers(scraper, DirName, ScraperDomain)

	// Population List Page
	scraper.OnXML("//table[@id='example2']/tbody/tr", func(element *colly.XMLElement) {
		country := element.ChildText("./td/a")
		url := element.ChildAttr("./td/a", "href")
		nextUrl := element.Request.AbsoluteURL(url)

		element.Request.Ctx.Put(nextUrl, country)

		if err := element.Request.Visit(nextUrl); err != nil {
			failedUrls[nextUrl] = err.Error()
			log.Println("Error visiting:", nextUrl)
		}
	})

	// Population Detail Page
	// scraper.OnHTML("table.table.table.table-condensed", func(h *colly.HTMLElement) {
	// 	fmt.Println("Here")
	// })
	scraper.OnXML("//table[@class='table table-striped table-bordered table-hover table-condensed table-list'][1]/tbody/tr", func(element *colly.XMLElement) {
		population := Population{
			Year:       element.ChildText("./td[1]"),
			Population: element.ChildText("./td[2]/strong"),
			Country:    element.Request.Ctx.Get(element.Request.URL.String()),
			Migrants:   element.ChildText("./td[5]"),
		}

		_ = writer.Write(population.ToSlice())
	})

	scraper.OnError(func(response *colly.Response, err error) {
		url := response.Request.URL.String()
		failedUrls[url] = err.Error()
	})

	scraperURL := fmt.Sprintf("https://%s/world-population/population-by-country/", ScraperDomain)

	if err := scraper.Visit(scraperURL); err != nil {
		log.Fatalf("Unable to start scraper:\n%v", err)
	}

	scraper.Wait()
}

package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func GetQuery() string {
	var query string
	for {
		fmt.Print("Please input your search query: ")
		query, _ = bufio.NewReader(os.Stdin).ReadString('\n')

		query = strings.TrimSpace(query)
		if !strings.EqualFold("", query) {
			break
		}
		fmt.Println("Query cannot be blank!")
	}

	return url.QueryEscape(query)
}

// OpenFile get file or create if not exist
func OpenFile(fileName string) *os.File {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln(err)
	}

	return file
}

// CloseFile close open file
func CloseFile(file *os.File) {
	err := file.Close()
	if err != nil {
		log.Println(err)
	}
}

func LogFailedUrls(dirName string, failedMap map[string]string) {
	if !(len(failedMap) > 0) {
		return
	}

	file := OpenFile(dirName + "failed-urls.txt")
	defer CloseFile(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	_ = writer.Write([]string{"URL", "ERROR"})

	for _url, err := range failedMap {
		log.Println(err)
		_ = writer.Write([]string{_url, err})
	}
}

func MakeDirs(dirName *string) {
	if !strings.HasSuffix(*dirName, "/") {
		*dirName = *dirName + "/"
	}

	if err := os.MkdirAll(*dirName+".cache/", os.ModePerm); err != nil {
		log.Fatalln(err)
	}
}

func RegisterHandlers(scraper *colly.Collector, dirName string, domain ...string) {
	scraper.CacheDir = dirName + ".cache/"
	scraper.AllowedDomains = append(
		scraper.AllowedDomains, domain...,
	)
	scraper.OnRequest(func(request *colly.Request) {
		log.Printf("Visiting: %s\n", request.URL.String())
	})
	scraper.OnResponse(func(response *colly.Response) {
		log.Printf(
			"Leaving: %s\nStatus Code: %d\n",
			response.Request.URL.String(), response.StatusCode,
		)
	})
}

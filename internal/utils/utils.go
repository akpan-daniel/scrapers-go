package utils

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
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

func GetFile() *os.File {
	var fileName string

	fmt.Print("Please enter name of output file (default: result): ")
	fileName, _ = bufio.NewReader(os.Stdin).ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	if strings.EqualFold(fileName, "") {
		fileName = "result"
	}

	fileName = strings.ReplaceAll(fileName, " ", "_")
	names := strings.Split(fileName, ".")
	if names[len(names)-1] != "csv" {
		fileName = strings.Join(names, "-") + ".csv"
	}

	return openFile(fileName)
}

// GetFile get file or create if not exist
func openFile(fileName string) *os.File {
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

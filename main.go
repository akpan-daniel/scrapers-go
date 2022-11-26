package main

import (
	"encoding/csv"
	"go-scrapers/internal/utils"
	"log"
)

func main() {
	query := utils.GetQuery()

	file := utils.GetFile()
	defer utils.CloseFile(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	log.Println(query)
}

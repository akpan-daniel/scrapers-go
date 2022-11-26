package main

import (
	"go-scrapers/internal/scrapers"
	"go-scrapers/internal/utils"
)

func main() {
	query := utils.GetQuery()

	scrapers.Init(query)
}

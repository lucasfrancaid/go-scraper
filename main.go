package main

import (
	"log"
	"os"

	"github.com/lucasfrancaid/go-scraper/scraper"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalln("Scraper argument in required, e.g: shopify")
	}
	scraperArg := os.Args[1]

	s, err := scraper.Set(scraperArg)
	if err != nil {
		log.Fatalln(err.Error())
	}

	s.Execute()
}

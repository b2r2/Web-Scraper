package main

import (
	"log"
	"os"
	"projects/Web-Scraper/utils"

	"github.com/anaskhan96/soup"
)

func main() {
	if len(os.Args) != 2 || !utils.IsURL(os.Args[1]) {
		log.Fatalf("Usage main.go http[s]://web-site...")
	}
	urlPath := os.Args[1]
	request, err := soup.Get(urlPath)
	if err != nil {
		log.Fatal(err)
	}
	//doc := soup.HTMLParse(request)
}

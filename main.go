package main

import (
	"log"
	"os"
	"projects/Web-Scraper/config"
	"projects/Web-Scraper/utils"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

var pattern map[string][]string = map[string][]string{
	"medium":  {"pre", "figure"},
	"telegra": {"pre", "figure"},
	"zen":     {"figure"},
}

func main() {
	if len(os.Args) != 2 || !utils.IsCorrectURL(os.Args[1]) {
		log.Fatalf("Missing URL argument. Usage main.go http[s]://web-site...")
	}
	urlPath := os.Args[1]
	domain := utils.GetDomain(urlPath)

	c := colly.NewCollector(colly.AllowURLRevisit())

	if domain == "telegra" {
		rp, err := proxy.RoundRobinProxySwitcher(config.Proxies...)
		if err != nil {
			log.Fatalf("Error when installing proxy ", err)
		}
		c.SetProxyFunc(rp)
	}
	c.Visit(urlPath)
}

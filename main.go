package main

import (
	"log"
	"os"
	"projects/Web-Scraper/config"
	"projects/Web-Scraper/utils"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

var patterns map[string][]string = map[string][]string{
	"medium":  {"section", "div"},
	"telegra": {"main", "header", "article"},
}

func main() {
	if len(os.Args) != 2 || !utils.IsCorrectURL(os.Args[1]) {
		log.Fatalf("Missing URL argument. Usage main.go http[s]://web-site...")
	}
	urlPath := os.Args[1]
	nameSite := utils.GetDomain(urlPath)

	c := colly.NewCollector(colly.AllowURLRevisit())
	// set proxy
	if nameSite == "telegra" {
		rp, err := proxy.RoundRobinProxySwitcher(config.Proxies...)
		if err != nil {
			log.Fatalf("Error when installing proxy ", err)
		}
		c.SetProxyFunc(rp)
	}
	// parsing content
	var contentPage []string
	var querySelector = patterns[nameSite][0]
	c.OnHTML(querySelector, func(e *colly.HTMLElement) {
		var temp string
		var tags []string = patterns[nameSite][1:]
		for _, tag := range tags {
			temp = e.ChildText(tag)
			contentPage = append(contentPage, temp)
		}
	})
	// set parallelism
	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})
	c.Visit(urlPath)
	c.Wait()
}

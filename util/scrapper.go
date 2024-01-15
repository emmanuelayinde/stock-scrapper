package util

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, price, change string
}

func Scrape(baseUrl string, tickers []string) []Stock {
	stocks := []Stock{}

	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Currently visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		stock := Stock{}

		stock.company = e.ChildText("h1")
		stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		stock.change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")

		stocks = append(stocks, stock)
	})

	c.Wait()

	for _, t := range tickers {
		c.Visit(baseUrl + t + "/")
	}

	return stocks
}

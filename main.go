package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type Stock struct {
	company, price, change string
}

func main() {
	baseUrl := "https://finance.yahoo.com/quote/"
	tickers := []string{
		"MSFT", "IBM", "GE", "UNP", "COST", "MCD", "V", "WMT", "DIS", "MMM", "INTC", "AXP", "AAPL", "BA", "CSCO", "GS", "JPM", "CRM", "VZ",
	}

	stocks := []Stock{}

	// Instantiate new colly
	c := colly.NewCollector()

	// Make request
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Currently visiting: ", r.URL)
	})

	// Listen for error
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	// Handle actually scrapping

	c.OnHTML("div#quote-header-info", func(e *colly.HTMLElement) {
		stock := Stock{}

		stock.company = e.ChildText("h1")
		stock.price = e.ChildText("fin-streamer[data-field='regularMarketPrice']")
		stock.change = e.ChildText("fin-streamer[data-field='regularMarketChangePercent']")

		stocks = append(stocks, stock)
	})

	c.Wait()

	// Loop through the available tickers
	for _, t := range tickers {
		c.Visit(baseUrl + t + "/")
	}

	fmt.Println(stocks)

	// Generate file
	file, err := os.Create("stocks.csv")
	if err != nil {
		log.Fatalln("Failed to create output CSV file: ", err)
	}
	defer file.Close()
	writer := csv.NewWriter(file)

	for _, stock := range stocks {
		record := []string{
			stock.company,
			stock.price,
			stock.change,
		}
		writer.Write(record)

	}
	defer writer.Flush()

}

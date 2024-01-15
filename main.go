package main

import (
	"time"

	"github.com/Emmanuelishola123/stock-scrapper/util"
)

func main() {
	baseUrl := "https://finance.yahoo.com/quote/"
	tickers := []string{
		"MSFT", "IBM", "GE", "UNP", "COST", "MCD", "V", "WMT", "DIS", "MMM", "INTC", "AXP", "AAPL", "BA", "CSCO", "GS", "JPM", "CRM", "VZ",
	}

	stocks := util.Scrape(baseUrl, tickers)

	// Generate file name based on current time and date
	currentTime := time.Now()
	fileName := "csv/stocks_" + currentTime.Format("2006-01-02_15-04-05") + ".csv"

	util.WriteCSV(fileName, stocks)
}

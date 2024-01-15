package util

import (
	"encoding/csv"
	"log"
	"os"
)

func WriteCSV(fileName string, stocks []Stock) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("Failed to create output CSV file: ", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, stock := range stocks {
		record := []string{
			stock.company,
			stock.price,
			stock.change,
		}
		writer.Write(record)
	}
}

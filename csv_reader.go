package main 

import (
	"encoding/csv"
	"log"
	"os"
)


func Reader() ([][]string, error) {
	problems, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("couldn't read the csv file", err)
	}

	reader := csv.NewReader(problems)
	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal("couldn't read the csv file", err)
	}

	return records, nil
}
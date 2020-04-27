package main

import (
	"encoding/csv"
	"log"
	"os"
)

var data = [][]string{{"Line1", "Hello Readers of"}, {"Line2", "golangcode.com"}, {"Line3", "Go 123"}}

func main() {
	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

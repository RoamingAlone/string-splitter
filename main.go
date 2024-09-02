package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func splitString(s string) (string, string) {
	openParen := strings.Index(s, "(")

	if openParen != -1 {
		beforeParen := strings.TrimSpace(s[:openParen])
		afterParen := strings.TrimSpace(s[openParen:]) // Include everything from the "(" onward
		return beforeParen, afterParen
	}

	return s, ""
}

func processCSV(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	header := records[0]
	header = append(header, "SKU", "Product Name")

	var newRecords [][]string
	newRecords = append(newRecords, header)

	for _, record := range records[1:] {
		var sku, productName string
		for _, colIndex := range []int{4, 5, 6, 7} { // columns E, F, G, H are indexes 4, 5, 6, 7
			part1, part2 := splitString(record[colIndex])
			if part2 != "" {
				sku = part1
				productName = part2
				break
			}
		}
		newRecord := append(record, sku, productName)
		newRecords = append(newRecords, newRecord)
	}

	outputFile, err := os.Create("output_modified.csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	for _, record := range newRecords {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	fileName := "QTY8.1-Table 1.csv" // replace with your file name
	if err := processCSV(fileName); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("CSV processing completed successfully.")
	}
}

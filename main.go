package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func splitString(s string) (string, string) {
	// Find the index of the opening parenthesis
	openParen := strings.Index(s, "(")
	if openParen == -1 {
		return s, ""
	}
	// Extract the part before the parenthesis
	beforeParen := strings.TrimSpace(s[:openParen])
	// Extract the part inside the parenthesis
	insideParen := strings.TrimSpace(s[openParen+1 : len(s)-1])
	return beforeParen, insideParen
}

func processCSV(fileName string, colIndex int) error {
	// Open the CSV file
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Process each record
	for i, record := range records {
		if i == 0 {
			record = append(record, "SKU", "Product Name")
			records[i] = record
			continue
		}
		if colIndex >= len(record) {
			return fmt.Errorf("Column index out of range")
		}

		// Split the content of the column
		beforeParen, insideParen := splitString(record[colIndex])
		// Add the split values to the record
		record = append(record, beforeParen, insideParen)
		records[i] = record
	}

	// Write the modified records to a new CSV file
	outputFile, err := os.Create("output.csv")
	if err != nil {
		return err
	}
	defer outputFile.Close()

	writer := csv.NewWriter(outputFile)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	fileName := "input.csv" //replace with actual name of the csv file
	colIndex := 1           // Replace with the index of the column you want to process

	if err := processCSV(fileName, colIndex); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("CSV processing completed successfully.")
	}
}

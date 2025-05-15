package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("raw_data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	lineCount := 0

	columnValues := make([][]string, 2)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)

		values := strings.Split(line, "  ")
		for index, value := range values {
			columnValues[index] = append(columnValues[index], value)
		}

		lineCount++
	}

	fmt.Println(columnValues)
	// handle head row
	headerValueMap := make(map[string][]string)
	for _, values := range columnValues {
		headerValueMap[values[0]] = values[1:]
	}

	fmt.Println(headerValueMap)

	result := make(map[string][]string)
	for _, tax := range headerValueMap["Tax Rates"] {
		result[tax] = []string{}
		for _, price := range headerValueMap["Prices"] {
			result[tax] = append(result[tax], price+tax)
		}
	}

	fmt.Println(result)
}

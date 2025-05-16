package main

import (
	"fmt"
	"github.com/kxw07/tax-calculator/file_ops"
)

func main() {
	columnValues, err := file_ops.Read("raw_data.txt")
	if err != nil {
		fmt.Println("Error when handling file:", err)
		return
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

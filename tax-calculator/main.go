package main

import (
	"fmt"
	"github.com/kxw07/tax-calculator/data_handle"
	"github.com/kxw07/tax-calculator/file_ops"
)

func main() {
	lines, err := file_ops.Read("raw_data.txt")
	if err != nil {
		fmt.Println("Error when handling file:", err)
		return
	}

	valuesByColumn := data_handle.SplitLinesByColumn(lines)

	valuesByHeader := data_handle.ExtractHeader(valuesByColumn)

	result := getTaxResult(valuesByHeader)

	fmt.Println(result)
}

func getTaxResult(valuesByHeader map[string][]int) map[int][]float64 {
	result := make(map[int][]float64)
	for _, tax := range valuesByHeader["Tax Rates"] {
		result[tax] = []float64{}
		for _, price := range valuesByHeader["Prices"] {
			result[tax] = append(result[tax], float64(price*(100+tax))/100)
		}
	}
	return result
}

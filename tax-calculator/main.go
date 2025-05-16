package main

import (
	"fmt"
	"github.com/kxw07/tax-calculator/file_ops"
	"strconv"
	"strings"
)

func main() {
	lines, err := file_ops.Read("raw_data.txt")
	if err != nil {
		fmt.Println("Error when handling file:", err)
		return
	}

	valuesByColumn := splitLinesByColumn(lines)

	valuesByHeader := extractHeader(valuesByColumn)

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

func extractHeader(dataMap [][]string) map[string][]int {
	headerValueMap := make(map[string][]int)
	for _, values := range dataMap {
		var ints []int
		for _, value := range values[1:] {
			intValue, _ := strconv.Atoi(value)
			ints = append(ints, intValue)
		}

		headerValueMap[values[0]] = ints
	}
	return headerValueMap
}

func splitLinesByColumn(lines []string) [][]string {
	columnValues := make([][]string, 2)
	lineCount := 0

	for _, line := range lines {
		values := strings.Split(line, "  ")
		for index, value := range values {
			columnValues[index] = append(columnValues[index], value)
		}

		lineCount++
	}

	return columnValues
}

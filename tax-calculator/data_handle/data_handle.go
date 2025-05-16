package data_handle

import (
	"strconv"
	"strings"
)

func ExtractHeader(dataMap [][]string) map[string][]float64 {
	headerValueMap := make(map[string][]float64)
	for _, columnValues := range dataMap {
		var valueArray []float64
		for _, columnValue := range columnValues[1:] {
			value, _ := strconv.ParseFloat(columnValue, 0)
			valueArray = append(valueArray, value)
		}

		headerValueMap[columnValues[0]] = valueArray
	}
	return headerValueMap
}

func SplitLinesByColumn(lines []string) [][]string {
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

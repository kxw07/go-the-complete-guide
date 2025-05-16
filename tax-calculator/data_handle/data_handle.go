package data_handle

import (
	"strconv"
	"strings"
)

func ExtractHeader(dataMap [][]string) map[string][]int {
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

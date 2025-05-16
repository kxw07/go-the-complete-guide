package file_ops

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Read(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
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

	return columnValues, nil
}

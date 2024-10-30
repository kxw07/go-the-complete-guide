package fileops

import (
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueString := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueString), 0644)
}

func ReadFloatFromFile(fileName string) float64 {
	valueBytes, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error reading from file")
		return 0
	}

	value, err := strconv.ParseFloat(string(valueBytes), 64)
	if err != nil {
		fmt.Println("Error parsing from file")
		return 0
	}

	return value
}

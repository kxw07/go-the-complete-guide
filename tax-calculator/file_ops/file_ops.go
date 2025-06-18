package file_ops

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type FileOps struct {
	InputPricesFilePath string
	OutputFilePath      string
}

func New(inputPricesFilePath, outputFilePath string) FileOps {
	return FileOps{
		InputPricesFilePath: inputPricesFilePath,
		OutputFilePath:      outputFilePath,
	}
}

func (fo FileOps) ReadPrices() ([]string, error) {
	file, err := os.Open(fo.InputPricesFilePath)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	prices := []string{}
	for scanner.Scan() {
		var price string
		_, err := fmt.Sscanf(scanner.Text(), "%v", &price)
		if err != nil {
			return nil, err
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (fo FileOps) WriteToFile(data interface{}) error {
	file, err := os.Create(fo.OutputFilePath)
	defer file.Close()

	if err != nil {
		return err
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return err
	}

	return nil
}

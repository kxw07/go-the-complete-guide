package file_ops

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type FileOps struct {
	InputTaxRatesFilePath string
	InputPricesFilePath   string
	OutputFilePath        string
}

func New(inputTaxRatesFilePath, inputPricesFilePath, outputFilePath string) *FileOps {
	return &FileOps{
		InputTaxRatesFilePath: inputTaxRatesFilePath,
		InputPricesFilePath:   inputPricesFilePath,
		OutputFilePath:        outputFilePath,
	}
}

func (fo FileOps) ReadPrices() ([]float64, error) {
	file, err := os.Open(fo.InputPricesFilePath)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	prices := []float64{}
	for scanner.Scan() {
		var price float64
		_, err := fmt.Sscanf(scanner.Text(), "%f", &price)
		if err != nil {
			return nil, err
		}
		prices = append(prices, price)
	}

	return prices, nil
}

func (fo FileOps) ReadTaxRates() ([]float64, error) {
	file, err := os.Open(fo.InputTaxRatesFilePath)
	defer file.Close()

	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	taxRates := []float64{}
	for scanner.Scan() {
		var taxRate float64
		_, err := fmt.Sscanf(scanner.Text(), "%f", &taxRate)
		if err != nil {
			return nil, err
		}
		taxRates = append(taxRates, taxRate)
	}

	return taxRates, nil
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

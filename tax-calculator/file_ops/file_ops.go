package file_ops

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func ReadPrices(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
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

func ReadTaxRates(fileName string) ([]float64, error) {
	file, err := os.Open(fileName)
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

func WriteToFile(fileName string, data interface{}) error {
	file, err := os.Create(fileName)
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

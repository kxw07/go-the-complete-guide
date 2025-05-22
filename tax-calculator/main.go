package main

import (
	"fmt"
	"github.com/kxw07/tax-calculator/file_ops"
	"github.com/kxw07/tax-calculator/product_info"
)

func main() {
	defer handlePanic()

	prices := getPrices()
	taxRates := getTaxRates()

	productInfo := product_info.NewProductInfo(taxRates, prices)
	productInfo.CalculatePricesAfterTax()

	fmt.Println(*productInfo)
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("catch panic:", r)
	}
}

func getTaxRates() []float64 {
	taxRates, err := file_ops.ReadTaxRates("tax_rates.txt")
	if err != nil {
		fmt.Println("Error when getTaxRates:", err)
		panic("Error when getTaxRates")
	}
	return taxRates
}

func getPrices() []float64 {
	prices, err := file_ops.ReadPrices("prices.txt")
	if err != nil {
		fmt.Println("Error when getPrices:", err)
		panic("Error when getPrices")
	}
	return prices
}

package main

import (
	"fmt"
	"github.com/kxw07/tax-calculator/file_ops"
	"github.com/kxw07/tax-calculator/product_info"
)

func main() {
	defer handlePanic()

	fileOps := file_ops.New("tax_rates.txt", "prices.txt", "prices_after_tax.txt")

	prices := getPrices(*fileOps)
	taxRates := getTaxRates(*fileOps)

	productInfo := product_info.NewProductInfo(taxRates, prices)
	productInfo.CalculatePricesAfterTax()
	writeToFile(*fileOps, productInfo)
}

func writeToFile(fileOps file_ops.FileOps, productInfo *product_info.ProductInfo) {
	err := fileOps.WriteToFile(productInfo)

	if err != nil {
		fmt.Println("Error when write to file:", err)
		panic("Error when write to file")
	}
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("catch panic:", r)
	}
}

func getTaxRates(fileOps file_ops.FileOps) []float64 {
	taxRates, err := fileOps.ReadTaxRates()
	if err != nil {
		fmt.Println("Error when getTaxRates:", err)
		panic("Error when getTaxRates")
	}
	return taxRates
}

func getPrices(fileOps file_ops.FileOps) []float64 {
	prices, err := fileOps.ReadPrices()
	if err != nil {
		fmt.Println("Error when getPrices:", err)
		panic("Error when getPrices")
	}
	return prices
}

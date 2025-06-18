package main

import (
	"fmt"
	"github.com/kxw07/tax-calculator/command_ops"
	"github.com/kxw07/tax-calculator/io_executor"
	"github.com/kxw07/tax-calculator/product_info"
)

func main() {
	defer handlePanic()

	taxRates := []string{"0.05", "0.1", "0.15"}

	for _, taxRate := range taxRates {
		//fileOps := file_ops.New("prices.txt", fmt.Sprintf("results_%s.json", taxRate))
		commandOps := command_ops.New()
		prices := getPrices(commandOps)

		productInfo := product_info.NewProductInfo(commandOps, taxRate, prices)
		productInfo.CalculatePricesAfterTax()
	}
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("catch panic:", r)
	}
}

func getPrices(iOExecutor io_executor.IOExecutor) []string {
	prices, err := iOExecutor.ReadPrices()
	if err != nil {
		fmt.Println("Error when getPrices:", err)
		panic("Error when getPrices")
	}
	return prices
}

package main

import (
	"fmt"
	"github.com/kxw07/tax-calculator/file_ops"
	"github.com/kxw07/tax-calculator/product_info"
)

func main() {
	taxRates := []string{"0.05", "0.1", "0.15"}

	for _, taxRate := range taxRates {
		fileOps := file_ops.New("prices.txt", fmt.Sprintf("results_%s.json", taxRate))
		//commandOps := command_ops.New()

		productInfo := product_info.New(fileOps, taxRate)

		err := productInfo.CalculatePricesAfterTax()
		if err != nil {
			fmt.Println("Error when CalculatePricesAfterTax:", err)
			panic("Error when CalculatePricesAfterTax")
		}
	}
}

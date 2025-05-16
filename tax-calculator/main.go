package main

import (
	"fmt"
	"github.com/kxw07/tax-calculator/data_handle"
	"github.com/kxw07/tax-calculator/file_ops"
	"github.com/kxw07/tax-calculator/product_info"
)

func main() {
	lines, err := file_ops.Read("raw_data.txt")
	if err != nil {
		fmt.Println("Error when handling file:", err)
		return
	}

	valuesByColumn := data_handle.SplitLinesByColumn(lines)

	valuesByHeader := data_handle.ExtractHeader(valuesByColumn)

	getTaxResult(valuesByHeader)
}

func getTaxResult(valuesByHeader map[string][]float64) {
	for _, taxRate := range valuesByHeader["Tax Rates"] {
		productInfo := product_info.NewProductInfo(taxRate, valuesByHeader["Prices"])
		productInfo.CalculatePricesAfterTax()
		fmt.Println(productInfo)
	}
}

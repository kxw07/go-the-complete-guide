package product_info

import (
	"fmt"
	"github.com/kxw07/tax-calculator/file_ops"
	"strconv"
)

type ProductInfo struct {
	FileOps file_ops.FileOps `json:"-"`
	TaxRate string           `json:"tax_rate"`
	Prices  []string         `json:"prices"`
}

func NewProductInfo(fileOps file_ops.FileOps, taxRate string, prices []string) *ProductInfo {
	return &ProductInfo{
		FileOps: fileOps,
		TaxRate: taxRate,
		Prices:  prices,
	}
}

func (productInfo *ProductInfo) CalculatePricesAfterTax() {
	taxPricesMap := make(map[string][]string)

	prices := make([]string, len(productInfo.Prices))
	for index, price := range productInfo.Prices {
		float64Price, _ := strconv.ParseFloat(price, 64)
		float64TaxRate, _ := strconv.ParseFloat(productInfo.TaxRate, 64)
		prices[index] = fmt.Sprintf("%.2f", float64Price*(1+float64TaxRate))
	}

	taxPricesMap[fmt.Sprintf("%v", productInfo.TaxRate)] = prices

	err := productInfo.FileOps.WriteToFile(taxPricesMap)
	if err != nil {
		fmt.Println("Error when write to file:", err)
		panic("Error when write to file")
	}
}

package product_info

import "fmt"

type ProductInfo struct {
	TaxRate        []float64           `json:"tax_rate"`
	Prices         []float64           `json:"prices"`
	PricesAfterTax map[string][]string `json:"prices_after_tax"`
}

func NewProductInfo(taxRate []float64, prices []float64) *ProductInfo {
	return &ProductInfo{
		TaxRate: taxRate,
		Prices:  prices,
	}
}

func (productInfo *ProductInfo) CalculatePricesAfterTax() {
	taxPricesMap := make(map[string][]string)

	for _, taxRate := range productInfo.TaxRate {
		prices := make([]string, len(productInfo.Prices))
		for index, price := range productInfo.Prices {
			prices[index] = fmt.Sprintf("%.2f", price*(1+taxRate))
		}

		taxPricesMap[fmt.Sprintf("%v", taxRate)] = prices
	}

	productInfo.PricesAfterTax = taxPricesMap
}

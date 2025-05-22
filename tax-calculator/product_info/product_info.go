package product_info

type ProductInfo struct {
	TaxRate        []float64             `json:"tax_rate"`
	Prices         []float64             `json:"prices"`
	PricesAfterTax map[float64][]float64 `json:"prices_after_tax"`
}

func NewProductInfo(taxRate []float64, prices []float64) *ProductInfo {
	return &ProductInfo{
		TaxRate: taxRate,
		Prices:  prices,
	}
}

func (productInfo *ProductInfo) CalculatePricesAfterTax() {
	taxPricesMap := make(map[float64][]float64)

	for _, taxRate := range productInfo.TaxRate {
		prices := make([]float64, len(productInfo.Prices))
		for _, price := range productInfo.Prices {
			prices = append(prices, price*(1+taxRate))
		}

		taxPricesMap[taxRate] = prices
	}

	productInfo.PricesAfterTax = taxPricesMap
}

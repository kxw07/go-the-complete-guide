package product_info

type productInfo struct {
	TaxRate        float64   `json:"tax_rate"`
	Prices         []float64 `json:"prices"`
	PricesAfterTax []float64 `json:"prices_after_tax"`
}

func NewProductInfo(taxRate float64, prices []float64) *productInfo {
	return &productInfo{
		TaxRate: taxRate,
		Prices:  prices,
	}
}

func (pricesAfterTax *productInfo) CalculatePricesAfterTax() {
	pricesAfterTaxList := make([]float64, len(pricesAfterTax.Prices))
	for i, price := range pricesAfterTax.Prices {
		pricesAfterTaxList[i] = price * (1 + pricesAfterTax.TaxRate)
	}

	pricesAfterTax.PricesAfterTax = pricesAfterTaxList
}

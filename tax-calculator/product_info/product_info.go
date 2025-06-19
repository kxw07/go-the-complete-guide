package product_info

import (
	"fmt"
	"github.com/kxw07/tax-calculator/io_executor"
	"strconv"
)

type ProductInfo struct {
	IOExecutor io_executor.IOExecutor `json:"-"`
	TaxRate    string                 `json:"tax_rate"`
	Prices     []string               `json:"prices"`
}

func New(io io_executor.IOExecutor, taxRate string) *ProductInfo {
	return &ProductInfo{
		IOExecutor: io,
		TaxRate:    taxRate,
	}
}

func (productInfo *ProductInfo) CalculatePricesAfterTax() error {
	err := productInfo.getPrices()
	if err != nil {
		return err
	}

	prices := make([]string, len(productInfo.Prices))
	for index, price := range productInfo.Prices {
		float64Price, _ := strconv.ParseFloat(price, 64)
		float64TaxRate, _ := strconv.ParseFloat(productInfo.TaxRate, 64)
		prices[index] = fmt.Sprintf("%.2f", float64Price*(1+float64TaxRate))
	}

	err = productInfo.IOExecutor.Write(prices)
	if err != nil {
		return err
	}

	return nil
}

func (productInfo *ProductInfo) getPrices() error {
	prices, err := productInfo.IOExecutor.ReadPrices()
	if err != nil {
		return err
	}

	productInfo.Prices = prices

	return nil
}

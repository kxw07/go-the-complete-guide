package io_executor

type IOExecutor interface {
	ReadPrices() ([]string, error)
	Write(data interface{}) error
}

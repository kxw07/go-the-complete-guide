package command_ops

import "fmt"

type CommandOps struct{}

func New() CommandOps {
	return CommandOps{}
}

func (commandOps CommandOps) ReadPrices() ([]string, error) {
	var prices []string

	for {
		var price string
		fmt.Println("Enter price (or 'exit' to finish): ")
		fmt.Scanf("%s", &price)

		if price == "exit" {
			break
		}

		prices = append(prices, price)
	}

	return prices, nil
}

func (commandOps CommandOps) Write(data interface{}) error {
	fmt.Println("Writing data to output:", data)
	return nil
}

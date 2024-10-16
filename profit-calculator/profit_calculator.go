package main

import (
	"fmt"
	"os"
)

func main() {
	revenue := getUserInput("Enter the revenue: ")
	expenses := getUserInput("Enter the expenses: ")
	tax := getUserInput("Enter the tax(%): ")

	ebt, profit, ratio := calculate(revenue, expenses, tax)

	fmt.Printf("The EBT is: %.2f\n", ebt)
	fmt.Printf("The PROFIT is: %.2f\n", profit)
	fmt.Printf("The RATIO is: %.2f\n", ratio)

	os.WriteFile("profit.txt", []byte(fmt.Sprintf("The EBT is: %.2f\nThe profit is: %.2f\nThe ratio is: %.2f\n", ebt, profit, ratio)), 0644)
}

func getUserInput(message string) float64 {
	var input float64

	fmt.Print(message)
	fmt.Scan(&input)

	if input <= 0 {
		panic("input must be greater than 0")
	}

	return input
}

func calculate(revenue float64, expenses float64, tax float64) (float64, float64, float64) {
	var ebt = revenue - expenses
	var profit = ebt * (1 - tax/100)
	var ratio = ebt / profit

	return ebt, profit, ratio
}

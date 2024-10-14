package main

import "fmt"

func main() {

	revenue, expenses, tax := getUserInput()

	ebt, profit, ratio := calculate(revenue, expenses, tax)

	fmt.Printf("The EBT is: %.2f\n", ebt)
	fmt.Printf("The profit is: %.2f\n", profit)
	fmt.Printf("The ratio is: %.2f\n", ratio)
}

func getUserInput() (float64, float64, float64) {
	var revenue float64
	var expenses float64
	var tax float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax(%): ")
	fmt.Scan(&tax)

	return revenue, expenses, tax
}

func calculate(revenue float64, expenses float64, tax float64) (float64, float64, float64) {
	var ebt = revenue - expenses
	var profit = ebt * (1 - tax/100)
	var ratio = ebt / profit

	return ebt, profit, ratio
}

package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var tax float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter the tax(%): ")
	fmt.Scan(&tax)

	var ebt = revenue - expenses
	var profit = ebt * (1 - tax/100)
	var ratio = ebt / profit
	fmt.Printf("The EBT is: %.2f\n", ebt)
	fmt.Printf("The profit is: %.2f\n", profit)
	fmt.Printf("The ratio is: %.2f\n", ratio)
}

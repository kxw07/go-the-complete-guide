package main

import "fmt"

func main() {
	// arrays
	prices := [3]float64{}
	prices[0] = 100.0
	prices[1] = 200.0
	prices[2] = 300.0
	fmt.Println(prices)

	// slices
	var numbers []int
	numbers = []int{1, 2, 3}
	fmt.Println(numbers)
	fmt.Println(numbers[:2])
	fmt.Println(numbers[1:])

	// slices is reference type
	slices := numbers[1:2]
	slices[0] = 25
	fmt.Println(slices)
	fmt.Println(numbers)
}

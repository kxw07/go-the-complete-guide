package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	doubleNumbers := transformNumbers(&numbers, doubleInt)
	fmt.Println(doubleNumbers)

	tripleNumbers := transformNumbers(&numbers, tripleInt)
	fmt.Println(tripleNumbers)
}

func transformNumbers(numbers *[]int, transform func(int) int) []int {
	result := []int{}

	for _, value := range *numbers {
		result = append(result, transform(value))
	}

	return result
}

func doubleInt(number int) int {
	return number * 2
}

func tripleInt(number int) int {
	return number * 3
}

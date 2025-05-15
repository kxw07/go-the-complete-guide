package main

import "fmt"

type transformFunc func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	anotherNumbers := []int{2, 3, 4, 5, 6}

	doubleNumbers := transformNumbers(&numbers, getTransformFunc(&anotherNumbers))
	fmt.Println("doubleNumbers", doubleNumbers)

	tripleNumbers := transformNumbers(&numbers, getTransformFunc(&numbers))
	fmt.Println("tripleNumbers", tripleNumbers)

	fmt.Println(variadicFunction(1, 2, 3, 4, 5))

	// better than above
	fmt.Println(variadicFunction(1, numbers...))
}

func variadicFunction(first int, numbers ...int) int {
	sum := first
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func transformNumbers(numbers *[]int, transform transformFunc) []int {
	result := []int{}

	for _, value := range *numbers {
		result = append(result, transform(value))
	}

	return result
}

func getTransformFunc(numbers *[]int) transformFunc {
	if (*numbers)[0]%2 == 0 {
		return double
	}

	return triple
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

package main

import "fmt"

func main() {
	age := 32
	var agePointer *int
	agePointer = &age

	fmt.Println("getAdultYears: ", getAdultYears(age))
	fmt.Println("getAdultYearsByPointer: ", getAdultYearsByPointer(agePointer))

	name := "John"
	var namePointer *string
	namePointer = &name

	fmt.Println("name: ", name)
	fmt.Println("namePointer: ", *namePointer)
}

func getAdultYears(age int) int {
	return age - 18
}

func getAdultYearsByPointer(age *int) int {
	return *age - 18
}

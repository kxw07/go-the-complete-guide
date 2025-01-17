package main

import "fmt"

func main() {
	// 1) Create a new array (!) that contains three hobbies you have
	// 		Output (print) that array in the command line.
	hobbies := [3]string{"reading", "coding", "gaming"}
	fmt.Println(hobbies)

	// 2) Also output more data about that array:
	//		- The first element (standalone)
	//		- The second and third element combined as a new list
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3) Create a slice based on the first element that contains
	//		the first and second elements.
	//		Create that slice in two different ways (i.e. create two slices in the end)
	hobbies_slice_one := hobbies[:2]
	hobbies_slice_two := []string{hobbies[0], hobbies[1]}
	fmt.Println(hobbies_slice_one)
	fmt.Println(hobbies_slice_two)

	// 4) Re-slice the slice from (3) and change it to contain the second
	//		and last element of the original array.
	hobbies_slice_one = hobbies_slice_one[1:3]
	fmt.Println(hobbies_slice_one)

	// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
	course_goals := []string{"learn go", "build a project"}
	fmt.Println(course_goals)

	// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
	course_goals[1] = "build a web app"
	course_goals = append(course_goals, "learn a new language")

	// 7) Bonus: Create a "Product" struct with title, id, price and create a
	//		dynamic list of products (at least 2 products).
	//		Then add a third product to the existing list of products.
	type product struct {
		title string
		id    int
		price float64
	}
	products := []product{
		{"book", 1, 10.99},
		{"laptop", 2, 999.99},
	}
	products = append(products, product{"phone", 3, 599.99})
	fmt.Println(products)

	// unpack list values
	origin := []int{1, 2, 3}
	should_be_append := []int{4, 5, 6}

	origin = append(origin, should_be_append...)
	fmt.Println(origin)
}

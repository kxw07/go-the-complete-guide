package main

import "fmt"

func main() {
	websites := map[string]string{
		"Google":   "https://www.google.com",
		"Facebook": "https://www.facebook.com",
	}

	websites["Youtube"] = "https://www.youtube.com"

	fmt.Println(websites)
	fmt.Println(websites["Google"])
	fmt.Println(websites["None"] == "")

	delete(websites, "Google")
	fmt.Println(websites)

	fruits := make([]string, 2, 4)
	fmt.Println(len(fruits), cap(fruits))
	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits = append(fruits, "Orange")
	fruits = append(fruits, "Pineapple")
	fruits = append(fruits, "Watermelon")
	fmt.Println(len(fruits), cap(fruits))
	fmt.Println(fruits)

	for key, value := range websites {
		fmt.Println("key:", key)
		fmt.Println("value:", value)
	}

	for index, value := range fruits {
		fmt.Println("index:", index)
		fmt.Println("value:", value)
	}

}

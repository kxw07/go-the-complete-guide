package main

import "fmt"

type User struct {
	firstName string
	lastName  string
	birthday  string
}

func main() {
	user := User{}
	user.firstName = getUserInput("Please enter your first name:")
	user.lastName = getUserInput("Please enter your last name:")
	user.birthday = getUserInput("Please enter your birthday:")

	fmt.Printf("Your full name is %s %s.\n", user.firstName, user.lastName)
	fmt.Printf("Your birthday is %s.\n", user.birthday)
}

func getUserInput(message string) string {
	fmt.Println(message)
	var value string
	fmt.Scanln(&value)

	return value
}

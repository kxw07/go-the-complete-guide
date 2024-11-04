package main

import (
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	birthday  string
	createdAt time.Time
}

func main() {
	// user := User{}
	// user.firstName = getUserInput("Please enter your first name:")
	// user.lastName = getUserInput("Please enter your last name:")
	// user.birthday = getUserInput("Please enter your birthday:")

	inputFirstName := getUserInput("Please enter your first name:")
	inputLastName := getUserInput("Please enter your last name:")
	inputBirthday := getUserInput("Please enter your birthday:")

	user := User{
		firstName: inputFirstName,
		lastName:  inputLastName,
		birthday:  inputBirthday,
		createdAt: time.Now(),
	}

	fmt.Printf("Your full name is %s %s.\n", user.firstName, user.lastName)
	fmt.Printf("Your birthday is %s.\n", user.birthday)
	fmt.Printf("Data was created at %s.\n", user.createdAt)
}

func getUserInput(message string) string {
	fmt.Println(message)
	var value string
	fmt.Scanln(&value)

	return value
}

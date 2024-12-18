package main

import (
	"fmt"

	"github.com/kxw07/structs/user"
)

func main() {
	// user := User{}
	// user.firstName = getUserInput("Please enter your first name:")
	// user.lastName = getUserInput("Please enter your last name:")
	// user.birthday = getUserInput("Please enter your birthday:")

	inputFirstName := getUserInput("Please enter your first name:")
	inputLastName := getUserInput("Please enter your last name:")
	inputBirthday := getUserInput("Please enter your birthday:")

	user := user.UserConstructor(inputFirstName, inputLastName, inputBirthday)

	// user := User{
	// firstName: inputFirstName,
	// lastName:  inputLastName,
	// birthday:  inputBirthday,
	// createdAt: time.Now(),
	// }

	user.OutputUserInfo()
	user.ClearName()
	user.OutputUserInfo()
}

func getUserInput(message string) string {
	fmt.Println(message)
	var value string
	fmt.Scanln(&value)

	return value
}

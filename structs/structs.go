package main

import (
	"fmt"

	"github.com/kxw07/structs/user"
)

func main() {
	inputFirstName := getUserInput("Please enter your first name:")
	inputLastName := getUserInput("Please enter your last name:")
	inputBirthday := getUserInput("Please enter your birthday:")

	user := user.Constructor(inputFirstName, inputLastName, inputBirthday)

	user.OutputUserInfo()
	user.ClearName()
	user.OutputUserInfo()
}

func getUserInput(message string) string {
	fmt.Print(message)
	var value string
	fmt.Scanln(&value)

	return value
}

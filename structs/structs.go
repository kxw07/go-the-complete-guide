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

func (u User) outputUserInfo() {
	fmt.Printf("Your full name is %s %s.\n", u.firstName, u.lastName)
	fmt.Printf("Your birthday is %s.\n", u.birthday)
	fmt.Printf("Data was created at %s.\n", u.createdAt)
}

func (u *User) clearName() {
	u.firstName = ""
	u.lastName = ""
}

func userConstructor(firstName string, lastName string, birthday string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}
}

func main() {
	// user := User{}
	// user.firstName = getUserInput("Please enter your first name:")
	// user.lastName = getUserInput("Please enter your last name:")
	// user.birthday = getUserInput("Please enter your birthday:")

	inputFirstName := getUserInput("Please enter your first name:")
	inputLastName := getUserInput("Please enter your last name:")
	inputBirthday := getUserInput("Please enter your birthday:")

	user := userConstructor(inputFirstName, inputLastName, inputBirthday)

	// user := User{
	// firstName: inputFirstName,
	// lastName:  inputLastName,
	// birthday:  inputBirthday,
	// createdAt: time.Now(),
	// }

	user.outputUserInfo()
	user.clearName()
	user.outputUserInfo()
}

func getUserInput(message string) string {
	fmt.Println(message)
	var value string
	fmt.Scanln(&value)

	return value
}

package user

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

func (u User) OutputUserInfo() {
	fmt.Printf("Your full name is %s %s.\n", u.firstName, u.lastName)
	fmt.Printf("Your birthday is %s.\n", u.birthday)
	fmt.Printf("Data was created at %s.\n", u.createdAt)
}

func (u *User) ClearName() {
	u.firstName = ""
	u.lastName = ""
}

func UserConstructor(firstName string, lastName string, birthday string) *User {
	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthday:  birthday,
		createdAt: time.Now(),
	}
}

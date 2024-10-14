package main

import "fmt"

func main() {
	var balance float64 = 1000
	var continued bool = true

	printWelcomeMessage()

	for continued {
		printChoices()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is: ", balance)
		} else if choice == 2 {
			amount := deposit()
			balance += amount
			fmt.Println("Your new balance is: ", balance)
		} else if choice == 3 {
			amount := withdraw()
			balance -= amount
			fmt.Println("Your new balance is: ", balance)
		} else if choice == 4 {
			continued = false
			fmt.Println("Bye bye!")
		} else {
			fmt.Println("Invalid choice")
		}
	}
}

func withdraw() float64 {
	fmt.Println("3. withdraw")
	fmt.Print("Enter the amount to withdraw: ")
	var amount float64
	fmt.Scan(&amount)

	return amount
}

func deposit() float64 {
	fmt.Println("2. deposit")
	fmt.Print("Enter the amount to deposit: ")
	var amount float64
	fmt.Scan(&amount)

	return amount
}

func printWelcomeMessage() {
	fmt.Println("Welcome to the bank")
}

func printChoices() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. check balance")
	fmt.Println("2. deposit")
	fmt.Println("3. withdraw")
	fmt.Println("4. exit")
}

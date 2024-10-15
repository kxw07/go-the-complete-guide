package main

import (
	"errors"
	"fmt"
)

func main() {
	var accountBalance float64 = 1000

	printWelcomeMessage()

	for {
		printChoices()

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Printf("Your balance is: %v\n\n", accountBalance)
		case 2:
			amount, err := deposit()

			if err != nil {
				fmt.Println(err)
				continue
			}

			accountBalance += amount
			fmt.Printf("Your new balance is: %v\n\n", accountBalance)
		case 3:
			amount, err := withdraw(accountBalance)

			if err != nil {
				fmt.Println(err)
				continue
			}

			accountBalance -= amount
			fmt.Printf("Your new balance is: %v\n\n", accountBalance)
		case 4:
			fmt.Println("Bye bye!")
			return
		default:
			fmt.Print("Invalid choice\n\n")
		}
	}
}

func withdraw(accountBalance float64) (float64, error) {
	fmt.Print("Enter the amount to withdraw: ")
	var amount float64
	fmt.Scan(&amount)

	if amount <= 0 {
		return amount, errors.New("invalid amount. must be greater than 0")
	}

	if amount > accountBalance {
		return amount, errors.New("invalid amount. must be greater than account balance")
	}

	return amount, nil
}

func deposit() (float64, error) {
	fmt.Print("Enter the amount to deposit: ")
	var amount float64
	fmt.Scan(&amount)

	if amount <= 0 {
		return amount, errors.New("invalid amount. must be greater than 0")
	}

	return amount, nil
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

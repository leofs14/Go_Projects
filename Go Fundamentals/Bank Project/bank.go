package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile) // _ In Go means a placeholder

	if err != nil { // Nil is when ou have no errors.
		return 1000, errors.New("failed to find balance file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return balance, nil
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint("Account Balance: ", balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err := getBalanceFromFile()

	if err != nil{
		fmt.Println("ERROR:")
		fmt.Println(err)
		fmt.Println("-----------")
		// Can use return here if you want to stop the aplication when you have an error
		// Also can use panic()
	}

	fmt.Println("Welcome to Go Bank!")

	for i := 0; i < 200; i++ {
		
		fmt.Println("What would you like to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your Current Balance is:", accountBalance)
		case 2:
			fmt.Print("Your Deposit: ")
			var depositAmmount float64
			fmt.Scan(&depositAmmount)

			if depositAmmount <= 0 {
				fmt.Println("Invalid deposit ammount.")
				//return // Cancel the program exec after error.
				continue // Skip the rest and return to the beggining
			}

			accountBalance += depositAmmount // Shortcut for 1 = 1+2
			fmt.Println("Balance Updated! New Amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 3:
			fmt.Print("Withdrawal Amount: ")
			var withdrawalAmmount float64
			fmt.Scan(&withdrawalAmmount)

			if withdrawalAmmount <= 0 || withdrawalAmmount > accountBalance {
				fmt.Println("Invalid withdrawal ammount.")
				//return // Cancel the program exec.
				continue // Part of the "for" loop, skip the rest of the code and return to the beginning
			}

			accountBalance -= withdrawalAmmount // Same Shortcut as above but for -
			fmt.Println("Balance Updated! New Amount:", accountBalance)
			writeBalanceToFile(accountBalance)
		case 4:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our bank!")
			//break // Exits the "for" loop. // Cant be used inside of switch because it breaks switch
			return // Exits the for loop because cant use break inside switch
		default: // When no other condition is met
			fmt.Println("Invalid Choice! Please choose one of the options available.")
			continue
		}
	}
}
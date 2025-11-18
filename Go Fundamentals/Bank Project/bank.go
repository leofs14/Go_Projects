package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil{
		fmt.Println("ERROR:")
		fmt.Println(err)
		fmt.Println("-----------")
		// Can use return here if you want to stop the aplication when you have an error
		// Also can use panic()
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7 at:",randomdata.PhoneNumber())

	for {	
		presentOptions()

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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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

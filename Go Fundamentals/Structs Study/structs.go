package main

import (
	"fmt"
	"example.com/structs/User"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser, err = user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Print(err)
		return // Exits the application
	}

	admin := user.NewAdmin("test@example.com", "password123")

	admin.OutputUserDetails()
	admin.ClearUserData()
	admin.OutputUserDetails()

	// ... do something awesome with that gathered data!

	appUser.OutputUserDetails() // Structs serve to group data to avoid conflicts with order or ...
	appUser.ClearUserData()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value) // Scan ln go to the next var when enter is pressed.
	return value
}
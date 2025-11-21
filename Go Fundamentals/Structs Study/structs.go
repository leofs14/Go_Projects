package main

import (
	"fmt"
	"time"
)

type User struct{ // Upper case because you may use in other packages.
	firstName string
	lastName string
	birthdate string
	createdAt time.Time // Struct imported, used in another struct
} // Declared out of the functions to be global.

func (u User) outputUserDetails () { // This is now a method of the struct User
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) clearUserData () { // Manipulating struct data.
	u.firstName = ""
	u.lastName = ""
} // Pointer is used because it edits the original and not a copy

func newUser (firstName, lastName, birthdate string) *User { // Struct Creator
	return &User{
		firstName: firstName,
		lastName: lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	} 
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser = newUser(userFirstName, userLastName, userBirthdate)

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails() // Structs serve to group data to avoid conflicts with order or ...
	appUser.clearUserData()
	appUser.outputUserDetails()
}



func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
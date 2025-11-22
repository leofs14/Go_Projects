package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct { // Upper case because you may use in other packages.
	firstName string
	lastName  string // Upper case so its accessible to other classes
	birthdate string
	createdAt time.Time // Struct imported, used in another struct
} // Declared out of the functions to be global.

type Admin struct {
	email string
	password string
	User // Upper case so other packages can access
}

func (u User) OutputUserDetails() { // This is now a method of the struct User
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserData() { // Manipulating struct data.
	u.firstName = ""
	u.lastName = ""
} // Pointer is used because it edits the original and not a copy

func NewAdmin(email, password string) Admin {
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "ADMIN",
			lastName: "ADMIN",
			birthdate: "----",
			createdAt: time.Now(),
		},
	}
}

func New(firstName, lastName, birthdate string) (*User, error) { // Struct Creator
	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("first name, last name and birthdate are required")
	} // Validating the constructor to make ir even more useful

	return &User{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}, nil
}
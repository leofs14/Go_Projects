package main //pointless to use pointers with small numbers like that, just learning in case i need it

import "fmt"

func main() {
	age := 32 //regular variable

	agePointer := &age // & is used when you want the address behind the value

	fmt.Println("Age:", *agePointer) // * is used when you want the func to use the value and not the address

	editAgeToAdultYears(agePointer)
	fmt.Println(age) // Change age value after casting the function
}

func editAgeToAdultYears(age *int){
	//return *age -18
	*age = *age -18
}
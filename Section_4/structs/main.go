package main

import "fmt"

type contactInfo struct {
	email string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// alex := person{"Alex", "Anderson"} // Define in this way, better is the next one
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)
	// var alex person // Alternative way to define struct
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Printf("%+v", alex)
	alex := person{
		firstName: "Alex",
		lastName: "Anderson",
		contactInfo: contactInfo{
			email: "alex@email.com",
			zipcode: 661166,
		},
	}

	alex.updateName("Alexi")
	alex.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
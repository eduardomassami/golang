package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type contactInfoTwo struct {
	complement string
	street     string
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
	contactInfoTwo
}

func main() {
	// alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	// jim := person{
	// 	firstName: "Jim",
	// 	lastName:  "Party",
	// 	contact: contactInfo{
	// 		email:   "jim@gmail.com",
	// 		zipCode: 94000,
	// 	},
	// 	contactInfoTwo: contactInfoTwo{
	// 		complement: "bloco16 apD21",
	// 		street:     "Lourdes Lopes Sanches",
	// 	},
	// }
	// fmt.Printf("%+v", jim)

	var jim person
	jim.firstName = "Jim"
	jim.lastName = "Party"
	jim.contact.email = "jim@gmail.com"
	jim.contact.zipCode = 94000
	jim.contactInfoTwo.complement = "bloco16 apD21"
	jim.contactInfoTwo.street = "Lourdes Lopes Sanches"

	// fmt.Printf("%+v", jim)
	// jim.print()

	// jimPointer := &jim
	// jimPointer.updateName("John")
	// jim.print()

	jim.updateName("John")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (pointerToPerson person) print() {
	fmt.Printf("%+v", pointerToPerson)
}

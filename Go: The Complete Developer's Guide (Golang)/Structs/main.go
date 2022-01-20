package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// 1st Method to declare structs
	// alex := person{firstName: "Alex", lastName: "Anderson"}
	// fmt.Println(alex)

	// 2nd Method to declard structs
	//var alex person
	//alex.firstName = "Alex"
	//alex.lastName = "Anderson"

	//fmt.Println(alex)
	//fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}
	// & will give the address of the value this variable is pointing at
	// * give the value from that memory address

	jim.updateName("John")
	// print result after update with pointer
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

// For define type positions *variable_name meaning is pointer type
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

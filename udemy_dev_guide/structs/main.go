package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	fmt.Printf("%+v", alex)

	joe := person{
		firstName: "Joe",
		lastName:  "String",
		contact: contactInfo{
			email:   "jestrahm@yahoo.com",
			zipCode: 9400,
		},
	}

	joePointer := &joe
	joePointer.updateName("Joseph")
	joe.print()

}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%v", p)
}

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

	// david := person{
	// 	firstName: "David",
	// 	lastName:  "Duarte",
	// }

	var david person

	david.firstName = "David"
	david.lastName = "Duarte"

	// david.contact = contactInfo{
	// 	email:   "davidandradeduarte@gmail.com",
	// 	zipCode: 123,
	// }

	david.contactInfo.email = "davidandradeduarte@gmail.com"
	david.contactInfo.zipCode = 213

	// david.print()

	// davidPointer := &david

	david.updateName("Alex")

	// fmt.Println(&david)
	david.print()

	// fmt.Println(david)
	// fmt.Printf("%+v", david)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

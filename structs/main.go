package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {

	robb := person{
		firstName: "Robb",
		lastName:  "Capistrano",
		contactInfo: contactInfo{
			email: "capistrano@capistrano.com",
			zip:   9231,
		},
	}

	robb.updateFirstName("Robert")

	robb.print()
}

func (pointerToPerson *person) updateFirstName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

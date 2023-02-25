package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email    string
	postcode string
}

func main() {
	p := person{
		firstName: "Neeraj",
		lastName:  "sidhaye",
		contactInfo: contactInfo{
			email:    "test@gmail.com",
			postcode: "m33",
		},
	}

	p.updateName("Sid")
	p.print()
}

func (p *person) updateName(name string) {
	(*p).firstName = name
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	contactInfo
}

type contactInfo struct {
	email  string
	mobile int
}

func main() {
	p := person{
		"neeraj",
		"sidhaye",
		contactInfo{
			"n@gmail.com",
			34353434,
		},
	}

	p.updatePerson()
	fmt.Println("updated person is", p)

}

func (p *person) updatePerson() {
	p.firstname = "James"
	p.lastname = "Mil"
}

package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	}
	p.lastName = spouseLastName
}

func main() {
	// Init person using struct
	person := Person{"Felipe", "Campara", "Dourados", "f", 26}
	// person1 := Person{firstName: "Felipe1", lastName: "Campara1", city: "Dourados1", gender: "m1", age: 261}

	// fmt.Println(person)
	// fmt.Println(person1)
	// person.age++
	// fmt.Println(person.firstName)
	// fmt.Println(person.greet())
	person.hasBirthday()
	person.getMarried("Will")
	fmt.Println(person)
}

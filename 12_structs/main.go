package main

import (
	"fmt"
	"strconv"
)

// Person Struct
type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	// age       int

	firstName, lastName, city, gender string
	age                               int
}

// Greeting Method (value reciever)
func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday Method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "M" {
		return
	}
	p.lastName = spouseLastName

}

func main() {
	// Init Person using Struct
	person1 := Person{firstName: "Devin", lastName: "Herath", city: "Kurunegala", gender: "M", age: 23}

	// Another way
	person2 := Person{"Devin", "Herath", "Kurunegala", "M", 23}
	person3 := Person{"Melinda", "Bush", "NY", "F", 30}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1)

	person1.hasBirthday()
	person1.hasBirthday()
	person1.getMarried("T")

	fmt.Println(person1.greet())

	person3.getMarried("Gates")
	fmt.Println(person3.greet())

}

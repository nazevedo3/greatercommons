// create a type person struct
// attach a method speak to type person using a pointer receiver
// *person
// create a type human interface
// to implicitly implement the interface, a human must have the speak method
// create func “saySomething”
// have it take in a human as a parameter
// have it call the speak method
// show the following in your code
// you CAN pass a value of type *person into saySomething
// you CANNOT pass a value of type person into saySomething

package main

import "fmt"

type person struct {
	first string
	last  string
}

func (p *person) speak() {
	fmt.Println(p.first, "says, my last name is", p.last)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Roger",
		last:  "Wilco",
	}

	p2 := person{
		first: "Sterling",
		last:  "Archer",
	}

	//This runs because the reciever for the speak method is a * to a person and I'm passing in the address of p1.
	saySomething(&p1)
	//This will NOT run because I'm passing in the the type person instead of a pointer to type person
	saySomething(p2)

	// 	//Receivers       Values

	// -----------------------------------------------

	// (t T)           T and *T

	// (t *T)          *T
}

// Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:

// first name
// last name
// favorite ice cream flavors
// Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

package main

import "fmt"

type person struct {
	first  string
	last   string
	flavor []string
}

func main() {

	p1 := person{
		first:  "James",
		last:   "Bond",
		flavor: []string{"Chocolate", "Strawberry"},
	}

	p2 := person{
		first:  "Miss",
		last:   "Moneypenny",
		flavor: []string{"Vanilla", "Peanut Butter"},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p2.first)
	fmt.Println(p2.last)

	for i, v := range p1.flavor {
		fmt.Println(i, v)
	}
}

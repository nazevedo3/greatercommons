// Take the code from the previous exercise, then store the values of type person in a map with the key of last name.
// Access each value in the map.
// Print out the values, ranging over the slice.

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

	x := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range x {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.flavor {
			fmt.Println(i, val)
		}
		fmt.Println("------")
	}
}

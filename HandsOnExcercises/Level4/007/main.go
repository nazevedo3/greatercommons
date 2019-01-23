// Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

// 		"James", "Bond", "Shaken, not stirred"

// 		"Miss", "Moneypenny", "Helloooooo, James."

// Range over the records, then range over the data in each record.

package main

import "fmt"

func main() {
	a := []string{
		"James",
		"Bond",
		"Shaken, not stirred",
	}

	b := []string{
		"Miss",
		"Moneypenny",
		"Hellooooo, James.",
	}

	x := [][]string{a, b}

	fmt.Printf("%T\n", x)

	for i, xs := range x {
		fmt.Println("record: ", i)
		for j, val := range xs {
			fmt.Printf("\t index position: %v \t value: \t %v \n", j, val)
		}
	}
}

// Using the code from the previous example, delete a record from your map.
// Now print the map out using the “range” loop

package main

import "fmt"

func main() {
	x := map[string][]string{
		`bond_james`: []string{`Shaken, not stirred`, `Martinis`, `Women`},

		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},

		`no_dr`: []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	x[`curry_steph`] = []string{`Basketball`, `Wife`, `Kids`}

	delete(x, `bond_james`)

	for k, v := range x {
		fmt.Println("The record for ", k)
		for i, val := range v {
			fmt.Printf("\t At index position %v, value is %v\n", i, val)
		}
	}
}

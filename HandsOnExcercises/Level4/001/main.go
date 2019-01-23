// Using a COMPOSITE LITERAL:
// create an ARRAY which holds 5 VALUES of TYPE int
// assign VALUES to each index position.
// Range over the array and print the values out.
// Using format printing
// print out the TYPE of the array

package main

import "fmt"

func main() {

	x := [5]int{10, 20, 30, 40, 50}
	for i, v := range x {
		fmt.Printf("At index positioin %v, the value is %v\n", i, v)
	}
	fmt.Printf("The type of x is %T\n", x)
}

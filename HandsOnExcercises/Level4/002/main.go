// Using a COMPOSITE LITERAL:
// create a SLICE of TYPE int
// assign 10 VALUES
// Range over the slice and print the values out.
// Using format printing
// print out the TYPE of the slice

package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for i, v := range x {
		fmt.Printf("In index position %v of the array, the value %v is stored.\n", i, v)
	}
	fmt.Printf("The type of x is %T\n", x)
}

//Create a for loop using this syntax

// for condition { }
// Have it print out the years you have been alive.

package main

import "fmt"

func main() {
	bd := 1983
	for bd <= 2019 {
		fmt.Println(bd)
		bd++
	}
}

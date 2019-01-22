// Building on the previous hands-on exercise, create a program that uses “else if” and “else”.

package main

import "fmt"

func main() {
	x := "James Bond"
	if x == "James Bond" {
		fmt.Println("Hello James")
	} else if x == "Miss Money Penny" {
		fmt.Println("Hello Miss Money Penny")
	} else {
		fmt.Println("I don't know you.")
	}

}

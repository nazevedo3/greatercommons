//Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {
	favSport := "basketball"
	switch favSport {
	case "baseball":
		fmt.Println("Baseball is my favorite sport")
	case "basketball":
		fmt.Println("Basketball is my favorite sport.")
	}
}

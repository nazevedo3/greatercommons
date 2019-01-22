//Create a program that uses a switch statement with no switch expression specified.

package main

import "fmt"

func main() {
	switch {
	case (2 == 1):
		fmt.Println("This should not print.")
	case (3 < 1):
		fmt.Println("This should not print.")
	case (100 == 100):
		fmt.Println("This should print.")
	}
}

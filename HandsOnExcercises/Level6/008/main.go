// Create a func which returns a func
// assign the returned func to a variable
// call the returned func

package main

import "fmt"

func main() {

	x := foo()
	fmt.Println(x())

}

func foo() func() string {
	return func() string {
		return "Hello from the func inside a func"
	}
}

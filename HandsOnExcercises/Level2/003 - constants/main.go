//Create TYPED and UNTYPED constants. Print the values of the constants.

package main

import "fmt"

const (
	a     = 43 //UNTYPED
	b int = 42 //TYPED
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
}

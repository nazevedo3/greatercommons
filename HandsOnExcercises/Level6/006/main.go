// Build and use an anonymous func

package main

import "fmt"

func main() {

	func() {
		fmt.Println("This ran from the anonymous func")
	}()

	fmt.Println("This ran from func main.")
}

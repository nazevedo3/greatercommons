// in addition to the main goroutine, launch two additional goroutines
// each additional goroutine should print something out
// use waitgroups to make sure each goroutine finishes before your program exists

package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("Being CPU", runtime.NumCPU())
	fmt.Println("Begin gs", runtime.NumGoroutine())

	wg.Add(2)

	go foo()
	go bar()
	wg.Wait()

	fmt.Println("End CPU", runtime.NumCPU())
	fmt.Println("End gs", runtime.NumGoroutine())
}

func foo() {
	fmt.Println("Hello from Foo!")
	wg.Done()
}

func bar() {
	fmt.Println("Hello from Bar!")
	wg.Done()

}

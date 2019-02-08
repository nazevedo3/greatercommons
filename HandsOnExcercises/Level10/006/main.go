// write a program that
// puts 100 numbers to a channel
// pull the numbers off the channel and print them

package main

import "fmt"

func main() {

	in := make(chan int)

	go send(in)

	for v := range in {
		fmt.Println(v)
	}

}

func send(a chan int) {

	for i := 0; i <= 10; i++ {
		a <- i
	}
	close(a)

}

// A “callback” is when we pass a func into a func as an argument. For this exercise,

// pass a func into a func as an argument

package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := sum(xi...)
	fmt.Println("The total sum of all the numbers is:", a)
	b := even(sum, xi...)
	fmt.Println("The total sum of the EVEN numbers is:", b)
}

func sum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v

	}
	return sum
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, val := range vi {
		if val%2 == 0 {
			yi = append(yi, val)
		}
	}
	return f(yi...)
}

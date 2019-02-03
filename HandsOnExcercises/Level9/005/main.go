// Fix the race condition you created in exercise #4 by using package atomic

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var wg sync.WaitGroup

	var counter int64
	gr := 50
	wg.Add(gr)

	for i := 0; i < gr; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
			wg.Done()

		}()
	}

	wg.Wait()
	fmt.Println("end value:", counter)
}

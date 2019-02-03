// Fix the race condition you created in the previous exercise by using a mutex

// it makes sense to remove runtime.Gosched()

package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	counter := 0
	gr := 50
	wg.Add(gr)

	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			m.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			m.Unlock()
			wg.Done()

		}()
	}

	wg.Wait()
	fmt.Println("end value:", counter)
}

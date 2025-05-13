package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func count(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		fmt.Printf("Goroutine %d: %d\n", id, i)
		time.Sleep(time.Duration(id) * 100 * time.Millisecond)
	}
}

func main() {
	var counter int64 = 0
	var wg sync.WaitGroup

	casIncrement := func(id int) {
		defer wg.Done()
		for j := 0; j < 10; j++ {
			for {
				old := atomic.LoadInt64(&counter)
				new := old + 1
				if atomic.CompareAndSwapInt64(&counter, old, new) {
					fmt.Printf("Goroutine %d: counter is %d\n", id, counter)
					break
				}
			}
			time.Sleep(100 * time.Millisecond)
		}
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go casIncrement(i)
	}

	wg.Wait()

	fmt.Println("Goodbye, World! - counter is ", counter)
}

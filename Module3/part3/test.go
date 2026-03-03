package part3

import (
	"fmt"
	"sync"
)

func Test() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var counter int // Our shared data

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			// This is the race condition!
			counter++
		}()
	}

	wg.Wait()
	fmt.Println("Final counter value:", counter)
}
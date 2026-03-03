package part4

import (
    "fmt"
    "sync"
)

func Prob1() {
    var wg sync.WaitGroup
    var counter int // Our shared data
		var mu sync.Mutex

    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            // This is the race condition!
						mu.Lock()
						defer mu.Unlock()
            counter++
        }()
    }

    wg.Wait()
    fmt.Println("Final counter value:", counter)
}
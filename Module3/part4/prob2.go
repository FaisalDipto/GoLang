//
// --- STARTING CODE ---
//
package part4

import (
    "fmt"
    "sync"
)

func Prob2() {
    var wg sync.WaitGroup
    // A map to store key-value pairs.
		var mu sync.Mutex
    concurrentMap := make(map[int]int)

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go func(n int) {
            defer wg.Done()
						mu.Lock()
						defer mu.Unlock()
            // This line causes a race condition!
            concurrentMap[n] = n * 2
        }(i)
    }

    wg.Wait()
    fmt.Println("Map size:", len(concurrentMap))
}
package part4

//
// --- STARTING CODE (BROKEN) ---
//

import (
    "fmt"
    "sync"
    // "time"
)

var mu sync.Mutex

// processData tries to process something, but might fail.
func processData(fail bool) {
    mu.Lock()
		defer mu.Unlock()

    if fail {
				// time.Sleep(time.Second * 2)
        fmt.Println("Operation failed, returning early.")
        return // <-- This is the problem area!
    }

    // This line is never reached if 'fail' is true.
    // mu.Unlock()
}

func Prob5() {
    processData(true) // This locks the mutex but never unlocks it.

    fmt.Println("Waiting for the lock...")

    // This second call will wait forever to get the lock.
    processData(false)

    fmt.Println("This line will never be printed.")
}
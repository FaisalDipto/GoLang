package part2

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %v done\n", id)
	time.Sleep(100 * time.Microsecond)
}

func Prob2(){
	var wg sync.WaitGroup
	wg.Add(1)
	go worker(1, &wg)
	wg.Add(1)
	go worker(2, &wg)
	wg.Add(1)
	go worker(3, &wg)
	wg.Wait()
}
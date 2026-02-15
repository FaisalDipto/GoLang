package part2

import (
	"fmt"
	"sync"
	"time"
)

func countdown(n int, wg *sync.WaitGroup){
	defer wg.Done()
	for i := n; i > 0; i--{
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}
}

func Prob1(){
	var wg sync.WaitGroup
	wg.Add(1)
	go countdown(5, &wg)
	fmt.Println("Starting Launch sequence in:")
	wg.Wait()
	fmt.Println("Liftoff")
}
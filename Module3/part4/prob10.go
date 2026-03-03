package part4

import (
	"fmt"
	"sync"
	"time"
)

func initializeConnection(){
	fmt.Println("Database connection initialized.")
	time.Sleep(time.Second)
}

func Prob10(){
	var on sync.Once
	var wg sync.WaitGroup

	for i:=0; i<10; i++{
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			on.Do(initializeConnection)
		}()
	}
	wg.Wait()
}
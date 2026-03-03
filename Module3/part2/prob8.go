package part2

import (
	"errors"
	"fmt"
	"sync"
)

type Result struct{
	Value int
	Err error
}

func safeDivide(a, b int, ch chan Result, wg *sync.WaitGroup){
	defer wg.Done()
	if b == 0 {
		ch <- Result{
			Value : 0,
			Err : errors.New("Error"),
		}
	}	else {
		ch <- Result{
			Value: a/b,
			Err: nil,
		}
	}
}

func Prob8(){
	var wg sync.WaitGroup
	ch := make(chan Result)
	wg.Add(1)
	go safeDivide(10, 2, ch, &wg)
	wg.Add(1)
	go safeDivide(5, 0, ch, &wg)
	go safeDivide(10, 2, ch, &wg)

	go func ()  {
		wg.Wait()
		close(ch)
	}()

	for c := range ch{
		result := c
		if result.Err != nil{
			fmt.Println(result.Err)
		}	else {
			fmt.Println(result.Value)
		}
	}
}
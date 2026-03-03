package part3

import (
	"fmt"
	"sync"
	"time"
)

func tryReceive(c chan string, wg *sync.WaitGroup){
	defer wg.Done()
	select{
	case res := <- c:
		fmt.Println(res)
	default:
		fmt.Println("No message available")
	}
}

func Prob4(){
	var wg sync.WaitGroup
	ch := make(chan string, 1)
	wg.Add(1)
	go tryReceive(ch, &wg)
	time.Sleep(time.Second * 2)
	ch <- "Sending message to the channel."
	wg.Add(1)
	go tryReceive(ch, &wg)

	wg.Wait()
	fmt.Println("Closing the program...")
}
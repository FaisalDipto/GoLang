package part3

import (
	"fmt"
	"time"
)

func Prob1(){
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func(){
		time.Sleep(time.Second * 2)
		ch1 <- "from worker 1"
	}()

	go func(){
		time.Sleep(time.Second * 1)
		ch1 <- "from worker 2"
	}()
	
	select{
	case res1 := <- ch1:
		fmt.Println(res1)
	case <- ch2:
		fmt.Println(ch2)
	}
}
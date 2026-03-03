package part3

import (
	"fmt"
	"time"
)

func Prob3(){
	ch := make(chan string)

	go func(){
		time.Sleep(time.Second * 3)
		ch <- "Task Completed!"
	}()

	select{
	case res := <- ch:
		fmt.Println(res)
	case <- time.After(time.Second * 4):
		fmt.Println("Error: Task timed out...")
	}
}
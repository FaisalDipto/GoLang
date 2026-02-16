package part2

import "fmt"

func produce(ch chan int){
	for i := 1; i <=5; i++{
		ch <- i
	}
	close(ch)
}

func Prob5(){
	ch := make(chan int)
	go produce(ch)

	for c := range ch {
		fmt.Println(c)
	}
}
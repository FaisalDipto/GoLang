package part2

import "fmt"

func Prob3(){
	ch := make(chan string)

	go func (){
		ch <- "Message received!"
	}()

	result := <- ch
	fmt.Println(result)
}
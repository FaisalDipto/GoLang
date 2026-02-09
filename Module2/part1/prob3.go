package part1

import "fmt"

func Prob3(){
	s := "Hello"
	sPtr := &s
	*sPtr = "GoodBye"
	fmt.Println(s)
}
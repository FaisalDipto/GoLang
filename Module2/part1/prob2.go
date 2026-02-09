package part1

import "fmt"

func Prob2(){
	x := 100
	p := &x
	fmt.Println(p)
	fmt.Println(*p)
}
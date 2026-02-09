package part1

import "fmt"

func resetCounter(counter *int){
	*counter = 0
}

func Prob6(){
	val := 15
	fmt.Println(val)
	resetCounter(&val)
	fmt.Println(val)
}
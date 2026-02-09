package part6

import "fmt"

func calculate(a, b int) (int, int){
	return a + b, a * b
}

func Prob4(){
	fmt.Println(calculate(5,6))
}
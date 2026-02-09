package part1

import "fmt"

func double(p *int){
	*p *= 2
}

func Prob4(){
	num := 5
	double(&num)
	fmt.Println(num)
}
package part6

import "fmt"

func sumAll(numbers ...int) int {
	sum := 0
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func Prob5(){
	fmt.Println(sumAll(1, 2, 3))
}
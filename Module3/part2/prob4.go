package part2

import "fmt"

func sumSLice(slice []int, ch chan int) {
	sum := 0
	for	_, slc := range slice{
		sum += slc
	}
	ch <- sum
}

func Prob4(){
	ch := make(chan int)
	slc := []int{1, 2, 3, 4, 5}
	go sumSLice(slc, ch)
	result := <- ch
	fmt.Println(result)
}
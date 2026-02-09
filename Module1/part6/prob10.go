package part6

import "fmt"

func calculateAverage(variadic ...float64) float64 {
	var sum float64
	length := len(variadic)
	if length == 0 {
		return 0
	}	else {
		for _, v := range variadic {
			sum += v
		}
		return sum / float64(length)
	}
}

func Prob10(){
	fmt.Println(calculateAverage(1.2,2.3,3.4))
}
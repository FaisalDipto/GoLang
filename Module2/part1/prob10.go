package part1

import "fmt"

func increment(counter *int){
	*counter += 1
}

func Prob10(){
	visits := 0
	for i := 1; i <= 5; i++ {
		increment(&visits)
	}
	fmt.Println(visits)
}
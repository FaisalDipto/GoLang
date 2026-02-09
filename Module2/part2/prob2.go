package part2

import "fmt"

func Prob2(){
	myCar := Car{
		Make: "BMW",
		Model: "GTX",
		Year: 1997,
	}

	fmt.Printf("%+v\n", myCar)
}
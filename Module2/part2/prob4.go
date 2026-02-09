package part2

import (
	"fmt"
	"math"
)



type Circle struct{
	Radius float64
}

func printCircle(c Circle){
	area := math.Pi * c.Radius * c.Radius
	fmt.Println(area)
}

func Prob4(){
	c := Circle{
		Radius: 5.6,
	}

	printCircle(c)
}
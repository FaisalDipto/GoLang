package part2

import "fmt"

type Student struct{
	Name string
	Grade int
}

func Prob3(){
	s1 := Student{
		// Name: "Faisal",
		Grade: 90,
	}

	s1.Grade = 100
	fmt.Println(s1.Grade)
}
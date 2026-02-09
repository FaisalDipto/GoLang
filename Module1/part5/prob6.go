package part5

import "fmt"

func Prob6(){
	score := 85
	if score >= 90 && score <= 100 {
		fmt.Println("A")
	} else if score >= 80 && score <= 89 {
		fmt.Println("B")
	} else if score >= 70 && score <= 79 {
		fmt.Println("C")
	} else if score >= 60 && score <= 69 {
		fmt.Println("D")
	}	else {
		fmt.Println("F")
	}
}
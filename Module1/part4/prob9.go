package part4

import "fmt"

func Prob9(){
	principal := 5000
	rate := 0.05
	timeInYears := 2.5

	interest := float64(principal) * rate * timeInYears
	
	fmt.Println(interest)
}
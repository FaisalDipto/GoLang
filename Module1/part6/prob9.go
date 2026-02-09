package part6

import "fmt"

func formatName(firstName, lastName string) string {
	return lastName + ", " + firstName
}

func Prob9(){
	fmt.Println(formatName("Faisal", "Dipto"))
}
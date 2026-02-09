package part5

import "fmt"

func Prob3(){
	day := "Wednesday"

	switch day {
	case "Monday":
		fmt.Println("Time to work!")
	case "Friday" :
		fmt.Println("Almost there!")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a regular day.")
	}
}
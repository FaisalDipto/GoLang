package part5

import "fmt"

func Prob2(){
	for i := 1; i <= 50; i++ {
		if i % 3 == 0 && i % 5 == 0 {
			fmt.Println("FizzBuzz")
		}	else if i % 3 == 0 {
			fmt.Println("Fizz")
		}	else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
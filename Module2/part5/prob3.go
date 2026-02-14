package part5

import "fmt"

func Prob3(){
	id := 123
	err := fmt.Errorf("Error: user with id %d not found!", id)

	fmt.Println(err)
}
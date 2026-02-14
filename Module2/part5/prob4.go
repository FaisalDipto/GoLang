package part5

import (
	"errors"
	"fmt"
)

func checkPositive(num int) error{
	if num < 0{
		return errors.New("Number must be positive.")
	}
	return nil
}

func Prob4(){
	err := checkPositive(1)
	if err != nil {
		fmt.Printf("error occured: %v\n", err)
	}

	err = checkPositive(-1)
	if err != nil {
		fmt.Printf("error occured: %v\n", err)
	}
}
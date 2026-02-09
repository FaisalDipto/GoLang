package part6

import (
	"errors"
	"fmt"
)

func safeDivide(numerator, denominator int) (int, error){
	if denominator == 0 {
		return 0, errors.New("Division by zero is not allowed.")
	}	else {
		return numerator/denominator, nil
	}
}

func Prob8(){
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}	else {
		fmt.Println(result)
	}
}
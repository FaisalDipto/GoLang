package part5

import (
	"errors"
	"fmt"
)

func validateAndDouble(num int) (int, error){
	if num <= 0 {
		return 0, errors.New("Number must be positive!")
	}
	return num * 2, nil
}

func Prob5(){
	result ,err := validateAndDouble(-1)

	if err != nil {
		fmt.Printf("Error occured: %v\n", err)
	} else {
		fmt.Println(result)
	}
}


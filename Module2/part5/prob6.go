package part5

import (
	"errors"
	"fmt"
)

func verifyAge(age int) error {
	if age <18{
		return errors.New("Age must be 18 or older")
	}
	return nil
}

func Prob6(){
	fmt.Println(verifyAge(15))
	fmt.Println(verifyAge(25))
}
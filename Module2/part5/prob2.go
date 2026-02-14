package part5

import (
	"errors"
	"fmt"
)

func Prob2(){
	error := errors.New("New error message")
	fmt.Println(error)
}
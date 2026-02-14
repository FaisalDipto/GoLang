package part5

import (
	"fmt"
	"strconv"
)

func Prob1(){
	val, err := strconv.Atoi("abc")
	if err != nil{
		fmt.Printf("Error occured: %v\n", err)
	} else {
		fmt.Println(val)
	}

	val, err = strconv.Atoi("123")
	if err != nil{
		fmt.Printf("Error occured: %v\n", err)
	} else {
		fmt.Println(val)
	}
}
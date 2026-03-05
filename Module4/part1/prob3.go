package part1

import (
	"fmt"
	"os"
)

func checkValue(v int) error {
	if v < 0{
		return fmt.Errorf("Invalid value: %d\n", v)
	}
	return nil
}

func validation(res error){
	if res == nil{
		fmt.Fprintf(os.Stdout, "Success!\n")
	}	else {
		fmt.Fprintf(os.Stderr, "Error occurred: %v.\n", res)
	}
}

func Prob3(){
	res := checkValue(5)
	validation(res)
	res = checkValue(-5)
	validation(res)
}
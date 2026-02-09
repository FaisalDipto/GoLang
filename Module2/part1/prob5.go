package part1

import "fmt"

var booleanVal = new(bool)

func Prob5(){
	fmt.Println(*booleanVal)
	*booleanVal = true
	fmt.Println(*booleanVal)

}

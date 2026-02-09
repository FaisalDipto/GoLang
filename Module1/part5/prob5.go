package part5

import "fmt"

var n int

func Prob5(){
	for {
		n += 1
		if n == 5 {
			fmt.Println("Breaking the loop")
			break
		}
		fmt.Println(n)
	}
}
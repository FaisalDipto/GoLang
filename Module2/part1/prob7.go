package part1

import "fmt"

func activateSubscription(status *bool){
	*status = true
}

func Prob7(){
	isActive := false
	activateSubscription(&isActive)
	fmt.Println(isActive)
}
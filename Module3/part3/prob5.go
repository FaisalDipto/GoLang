package part3

import "fmt"

func trySend(c chan int, val int){
	select{
	case c <- val:
		fmt.Println("Sent message.")
	default:
		fmt.Println("Channel is full. Could not send.")
	}
}

func Prob5(){
	ch := make(chan int, 1)
	trySend(ch, 1)
	trySend(ch, 2)
}
package part1

import "fmt"

func updatePort(portPtr *int){
	*portPtr = 8080
}

func Prob9(){
	port := 80
	updatePort(&port)
	fmt.Println(port)
}
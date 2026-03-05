package part1

import (
	"fmt"
	"time"
)

func prog()(string, int){
	var name string
	var year int
	fmt.Println("Enter your name")
	fmt.Scanln(&name)
	fmt.Println("Enter your birth year")
	fmt.Scanln(&year)
	
	age := time.Now().Year() - year

	return name, age
}

func Prob2(){
	name, age := prog()
	fmt.Printf("Hello %v, you are approximately %v years old\n", name, age)
}



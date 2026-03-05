package part1

import "fmt"

type User struct{
	ID int
	Name string
}

func Prob1(){
	u1 := User{
		ID: 1,
		Name: "Faisal",
	}
	fmt.Printf("User : %v\n", u1)
	fmt.Printf("User : %+v\n", u1)
	fmt.Printf("User : %#v\n", u1)
}
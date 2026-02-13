package part4

import "fmt"

type Dog struct{
	Name string
}

func (d Dog) Speak() string{
	return "Woof! My name is " + d.Name
}

func Prob3(){
	d1 := Dog{
		Name: "Butch",
	}

	fmt.Println(d1.Speak())
}
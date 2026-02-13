package part4

import "fmt"

type Person struct{
	Name string
}

func (p Person) Speak() string {
	return "Hello, my name is " + p.Name
}

func Prob2(){
	p1 := Person{
		Name: "Faisal",
	}

	fmt.Println(p1.Speak())
}
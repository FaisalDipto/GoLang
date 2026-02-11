package part3

import "fmt"

type Greeter struct{
	Greeting string
}

func (g Greeter) Greet(name string) string {
	return g.Greeting + ", " + name
}

func Prob2(){
	g1 := Greeter{
		Greeting: "Hello",
	}

	msg := g1.Greet("Faisal")
	fmt.Println(msg)
}
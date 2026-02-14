package part4

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Value: %+v, Type: %T\n", i, i)
}

func Prob10(){
	integer := 1
	strg := "Hello, World!"
	p5 := Person{
		Name: "Amir",
	}

	describe(integer)
	describe(strg)
	describe(p5)
}
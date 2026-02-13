package part4

import "fmt"

func makeSound(s Speaker){
	fmt.Println(s.Speak())
}

var P2 = Person{
	Name: "Dipto",
}

var D2 = Dog{
	Name: "Bolt",
}

func Prob4(){
	// P2 := Person{
	// 	Name: "Dipto",
	// }

	// D2 := Dog{
	// 	Name: "Bolt",
	// }

	makeSound(P2)
	makeSound(D2)
}
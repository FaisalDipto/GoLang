package main

import (
	"fmt"
	"mycalculator/mathops"

	"rsc.io/quote"
)

func main(){
	sum := mathops.Add(1, 2)
	fmt.Println("The sum is:", sum)

	fmt.Println("A classic Go proverb:", quote.Go())
}
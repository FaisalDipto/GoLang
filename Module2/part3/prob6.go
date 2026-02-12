package part3

import "fmt"

type Car struct{
	IsEngineOn bool
}

func (c *Car) StartEngine(){
	c.IsEngineOn = true
}

func (c *Car) StopEngine(){
	c.IsEngineOn = false
}

func Prob6(){
	c1 := Car{}
	fmt.Println(c1.IsEngineOn)
	c1.StartEngine()
	fmt.Println(c1.IsEngineOn)
	c1.StopEngine()
	fmt.Println(c1.IsEngineOn)
}
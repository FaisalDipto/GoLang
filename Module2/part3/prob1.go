package part3

import "fmt"

type Rectangle struct{
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func Prob1(){
	r1 := Rectangle{
		Width: 12.3,
		Height: 15.6,
	}

	fmt.Println(r1.Area())
}
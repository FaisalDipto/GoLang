package part6

import "fmt"

func calculateRectangleArea(width, height float64) float64 {
	return width * height
}

func Prob6(){
	fmt.Println(calculateRectangleArea(5, 6))
}
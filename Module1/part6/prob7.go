package part6

import "fmt"

func convertCelsiusToFahrenheit(f float64) float64{
	return (f * 9/5) + 32
}

func Prob7(){
	fmt.Println(convertCelsiusToFahrenheit(12))
}
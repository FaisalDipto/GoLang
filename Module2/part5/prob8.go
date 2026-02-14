package part5

import (
	"errors"
	"fmt"
	"math"
)

func safeSqrt(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("Cannot calculate square with a negative number")
	}
	return math.Sqrt(num), nil
}

func Prob8(){
	fmt.Println(safeSqrt(9))
	fmt.Println(safeSqrt(-9))
}
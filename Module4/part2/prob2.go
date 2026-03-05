package part2

import (
	"fmt"
	"strconv"
	"strings"
	"os"
)

func Prob2(){
	numbers := []int{10, 20, 30, 40, 50}
	var numList []string


	for _, number := range numbers{
		num := strconv.Itoa(number)
		numList = append(numList, num)
	}

	strg := fmt.Sprintf("Report Data: [%v]\n",strings.Join(numList, ","))
	fmt.Fprintf(os.Stdout, "%v\n", strg)
}
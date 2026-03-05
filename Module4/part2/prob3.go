package part2

import (
	"fmt"
	"strconv"
	"strings"
)

func Prob3(){
	query := "id=123&user=faisal&active=true&status=online"
	pairs := strings.Split(query, "&")
	for _, pair := range pairs{
		list := strings.Split(pair, "=")
		key := list[0]
		value := list[1]
		switch key {
		case "id":
			valInt, err := strconv.Atoi(value)
			if err != nil{
				fmt.Printf("Error occurred %v\n: ", err)
			}	else {
				fmt.Printf("The key is %v & the value is %v, type %T\n", key, valInt, valInt)
			}

		case "active":
			valBool, err := strconv.ParseBool(value)
			if err != nil{
				fmt.Printf("Error occurred %v\n: ", err)
			}	else {
				fmt.Printf("The key is %v & the value is %v, type %T\n", key, valBool, valBool)
			}

		default:
			fmt.Printf("The key is %v & the value is %v, type %T\n", key, value, value)
		}
	}
}
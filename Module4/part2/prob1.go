package part2

import (
	"fmt"
	// "strconv"
	"strings"
)

func Prob1(){
	data := "		user	,		admin,guest,	super	user"
	dataSlice := strings.Split(data, ",")
	for _, dataFin := range dataSlice{
		fmt.Println(strings.TrimSpace(dataFin))
	}
}
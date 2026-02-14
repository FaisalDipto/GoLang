package part5

import (
	"errors"
	"fmt"
)

func validateUsername(username string) error{
	if username == "" || len(username) > 20 {
		return errors.New("Username invalid")
	}
	return nil
}

func Prob7(){
	fmt.Println(validateUsername("Faisal"))
	fmt.Println(validateUsername(""))
	fmt.Println(validateUsername("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
}
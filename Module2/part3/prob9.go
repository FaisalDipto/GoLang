package part3

import (
	"fmt"
	"strings"
)

type User struct{
	Email string
}

func (u User) IsValidEmail() bool {
	return strings.Contains(u.Email, "@")
}

func Prob9(){
	u1 := User{
		Email: "dipto@gmail.com",
	}

	u2 := User{
		Email: "faisalgmail.com",
	}

	fmt.Println(u1.IsValidEmail())
	fmt.Println(u2.IsValidEmail())
}


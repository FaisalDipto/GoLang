package part2

import "fmt"

type UserProfile struct{
	Username string
	FollowerCount int
	IsVarified bool
}

func Prob6(){
	u1 := UserProfile{
		Username: "Faisal",
		FollowerCount: 0,
		IsVarified: false,
	}

	fmt.Printf("%+v", u1)
}
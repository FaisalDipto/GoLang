package part3

import "fmt"

type User struct{
	FirstName string
	LastName string
}

func (u User) FullName() string{
	return u.FirstName + ", " + u.LastName
}

func Prob5(){
	u1 := User{
		FirstName: "Faisal",
		LastName: "Dipto",
	}

	fmt.Println(u1.FullName())
}
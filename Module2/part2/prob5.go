package part2

import "fmt"

type User struct{
	Email string
}

func updateEmail(u *User, newEmail string){
	u.Email = newEmail
}

func Prob5(){
	u1 := User{
		Email: "faisal@gmail.com",
	}

	newEmail := "dipto@gmail.com"
	updateEmail(&u1, newEmail)
	fmt.Println(u1.Email)
}
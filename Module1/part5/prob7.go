package part5

import "fmt"

func Prob7(){
	saveUsername := "Faisal"
	savedPassword := "1234"
	
	enteredUsername := "Dipto"
	enteredPassword := "4567"

	if saveUsername == enteredUsername && savedPassword == enteredPassword {
		fmt.Println("Login Successful")
	}	else {
		fmt.Println("Invalid Credentials")
	}
}
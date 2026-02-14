package part5

import "fmt"

func checkPasswordStrength(password string) error{
	if len(password) < 8{
		return fmt.Errorf("Password is too short: must be at least 8 characters, but got %d\n", len(password))
	}
	return nil
}

func Prob10(){
	fmt.Println(checkPasswordStrength("1234"))
}
package part5

import (
	"fmt"
)

func openFile(filename string) error{
	if filename == "data.csv"{
		return nil
	}
	return fmt.Errorf("File not found: '%s'", filename)
}

func Prob9(){
	fmt.Println(openFile("data.csv"))
	fmt.Println(openFile("passwords.txt"))
}
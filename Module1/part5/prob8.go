package part5

import "fmt"

func Prob8(){
	command := "start"

	switch command {
	case "start":
		fmt.Println("Starting service...")
	case "stop":
		fmt.Println("Stopping service...")
	case "restart":
		fmt.Println("Restarting service...")
	default:
		fmt.Println("Unknown command.")
	}
}
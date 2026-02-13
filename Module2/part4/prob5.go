package part4

import "fmt"

func Prob5(){
	// slcSpeakers := []Speaker{
	// 	Person{
	// 		Name: "Faisal",
	// 	},
	// 	Dog{
	// 		Name: "Butch",
	// 	},
	// }

	// for _, slc := range slcSpeakers{
	// 	fmt.Println(slc.Speak())
	// }

	slcSpeakers := []Speaker{
		P2, D2,
	}

	for _, slc := range slcSpeakers{
		fmt.Println(slc.Speak())
	}
}
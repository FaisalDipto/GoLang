package part5

import "fmt"


func Prob10(){
	sentence := "The quick brown fox jumps over the lazy dog"

	count := 1

	for _, s := range sentence {
		switch s{
		case 'a', 'e', 'i', 'o', 'u':
			count += 1
		}
	}

	fmt.Println(count)
}
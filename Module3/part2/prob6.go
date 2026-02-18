package part2

import (
	"fmt"
	"time"
)

func checkWebsite(url string, ch chan string) {
	time.Sleep(50 * time.Millisecond)
	str := url + " is OK"
	ch <- str
}

func Prob6(){
	strSlc := []string{
		"google.com",
		"amazon.com",
		"github.com",
	}

	ch := make(chan string)

	for _, str := range strSlc{
		go checkWebsite(str, ch)
	}

	for i := 0; i < len(strSlc); i++ {
		c := <-ch
		fmt.Println(c)
	}
}
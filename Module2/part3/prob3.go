package part3

import "fmt"

type Counter struct {
	Value int
}

func (c *Counter) Increment() {
	c.Value += 1
}

func Prob3(){
	c1 := Counter{
		Value: 1,
	}

	c1.Increment()
	fmt.Println(c1.Value)
}
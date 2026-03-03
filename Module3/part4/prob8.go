package part4

import (
	"fmt"
	"strconv"
	"sync"
)

type Logger struct{
	messages []string
	mu sync.Mutex
}

func (l *Logger) Log(message string){
	l.mu.Lock()
	defer l.mu.Unlock()
	l.messages = append(l.messages, message)
}

func Prob8(){
	l1 := Logger{}
	var wg sync.WaitGroup

	for i:= 0; i<100; i++{
		wg.Add(1)
		go func (i int)  {
			defer wg.Done()
			for j:= 0; j<5; j++{
				id1 := strconv.Itoa(i)
				id2 := strconv.Itoa(j)
				fullId := id1 + id2
				l1.Log(fullId)
			}
		}(i)
	}
	
	wg.Wait()
	fmt.Println(len(l1.messages))
}
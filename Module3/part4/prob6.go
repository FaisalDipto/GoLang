package part4

import (
	"fmt"
	"sync"
)

type ShoppingCart struct{
	IDs map[string]int
	mu sync.Mutex
}

func (s *ShoppingCart) AddItem(itemID string){
	s.mu.Lock()
	defer s.mu.Unlock()
	s.IDs[itemID] += 1
}

func Prob6(){
	var wg sync.WaitGroup
	s1 := ShoppingCart{
		IDs : make(map[string]int),
	}
	for i := 0; i<100; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			s1.AddItem("TV")
		}()
	}

	wg.Wait()
	fmt.Println(s1.IDs["TV"])
}
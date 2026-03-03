package part4

import (
	"fmt"
	"sync"
)

type Player struct{
	X int
	Y int
	mu sync.Mutex
}

func (p *Player) Move(dx, dy int) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.X += dx
	p.Y += dy
}

func Prob9(){
	p1 := Player{
		X: 0,
		Y: 0,
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func ()  {
		defer wg.Done()
		for i:= 0; i<1000; i++{
			p1.Move(1,1)
		}
	}()

	wg.Add(1)
	go func () {
		defer wg.Done()
		for i:=0; i<1000; i++{
			p1.Move(1,1)
		}
	}()
	wg.Wait()
	fmt.Println(p1.X, p1.Y)
}
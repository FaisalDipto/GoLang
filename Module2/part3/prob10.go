package part3

import "fmt"

type Player struct{
	Name string
	Health int
}

func (p *Player) TakeDamage(amount int) {
	if p.Health >= amount {
		p.Health -= amount
	}	else {
		p.Health = 0
	}
}

func Prob10(){
	p1 := Player{
		Health: 100,
	}

	p1.TakeDamage(40)
	p1.TakeDamage(70)

	fmt.Println(p1.Health)
}
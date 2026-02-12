package part3

import "fmt"

func (p *Product) ApplyDiscount(f float64) {
	p.Price = p.Price - p.Price * (f/100)
}

func Prob8(){
	p2 := Product{
		Price: 100,
	}

	fmt.Println(p2.Price)
	p2.ApplyDiscount(10)
	fmt.Println(p2.Price)
}
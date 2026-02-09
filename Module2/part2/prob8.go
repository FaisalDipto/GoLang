package part2

import "fmt"

func applyDiscount(p *Product, discount float64){
	discountedPrice := p.Price - (p.Price * (discount/100))
	p.Price = discountedPrice
}

func Prob8(){
	p1 := Product{
		ID: "145",
		Name: "Notebook",
		Price: 100,
	}

	originalPrice := p1.Price
	fmt.Println(originalPrice)

	applyDiscount(&p1, 10.0)
	fmt.Println(p1.Price)
}
package part4

import "fmt"

type Product struct{
	Name string
	Price float64
}

func (p Product) String() string{
	return fmt.Sprintf("Product: %v ($%.2f)\n", p.Name, p.Price)
}

func Prob8(){
	p3 := Product{
		Name: "TV",
		Price: 1000.00,
	}

	fmt.Println(p3)
}

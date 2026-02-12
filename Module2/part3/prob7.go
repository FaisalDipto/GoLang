package part3

import "fmt"

type Product struct{
	Name string
	Price float64
}

func (p Product) PrintDetails(){
	fmt.Printf("Product: %v, Price: $%.2f", p.Name, p.Price)
}

func Prob7(){
	p1 := Product{
		Name: "Laptop",
		Price: 1200.00,
	}

	p1.PrintDetails()
}
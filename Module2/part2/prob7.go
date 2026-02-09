package part2

import "fmt"

type Product struct{
	ID string
	Name string
	Price float64
}

func Prob7(){
	p1 := Product{
		ID: "123",
		Name: "Book",
		Price: 700,
	}

	p2 := Product{
		ID: "124",
		Name: "TV",
		Price: 15000,
	}
	Products := []Product{p1, p2}
	fmt.Printf("%+v", Products)
}
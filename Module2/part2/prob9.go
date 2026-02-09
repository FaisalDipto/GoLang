package part2

import "fmt"

type Address struct{
	Street string
	City string
	Zipcode string
}

type Contact struct{
	Name string
	Phone string
	Address
}

func Prob9(){
	c1 := Contact{
		Name: "Faisal",
		// Phone: "12345",
		Address: Address{
			Street: "Mirpur",
			City: "Dhaka",
			Zipcode: "1216",
		},
	}

	fmt.Println(c1.Name, c1.City)
}
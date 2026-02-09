package part2

import "fmt"

type Employee struct{
	Name string
	Salary float64
}

func isHigherPaid(e1 Employee, e2 Employee) bool {
	if e1.Salary > e2.Salary {
		return true
	}	else {
		return false
	}
}

func Prob10(){
	e1 := Employee{
		Name: "Faisal",
		Salary: 60000,
	}

	e2 := Employee{
		Name: "Faisal",
		Salary: 80000,
	}

	result := isHigherPaid(e2, e1)
	fmt.Println(result)
}
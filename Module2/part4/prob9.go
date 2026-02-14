package part4

import "fmt"

type PaymentMethod interface{
	Pay(amount float64) string
}

type CreditCard struct{}
type Paypal struct{}

func (c CreditCard) Pay(amount float64) string{
	return fmt.Sprintf("Paid $%.2f using Credit Card.\n", amount)
}

func (p Paypal) Pay(amount float64) string{
	return fmt.Sprintf("Paid $%.2f using Paypal.\n", amount)
}

func Prob9(){
	lists := []PaymentMethod{
		CreditCard{}, Paypal{}}
	
	for _, list := range lists{
		fmt.Println(list.Pay(100.00))
	}
}



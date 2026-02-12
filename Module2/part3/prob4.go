package part3

import (
	"fmt"
)

type Account struct{
	Balance float64
}

func (a *Account) AddBalancePointer(amount float64) {
	a.Balance += amount
}

func (a Account) AddBalanceValue(amount float64) {
	a.Balance += amount
}

func Prob4(){
	a1 := Account{
		Balance: 500.99,
	}

	a1.AddBalanceValue(20.00)
	fmt.Println(a1.Balance)
	a1.AddBalancePointer(20.00)
	fmt.Println(a1.Balance)
}


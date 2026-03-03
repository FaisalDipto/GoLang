package part4

import (
	"fmt"
	"sync"
)

type BankAccount struct{
	balance int
	mu sync.Mutex
}

func (b *BankAccount) Deposit(amount int){
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance += amount
}
func (b *BankAccount) Withdraw(amount int){
	b.mu.Lock()
	defer b.mu.Unlock()
	b.balance -= amount
}

func Prob3(){
	var wg sync.WaitGroup
	b1 := BankAccount{
	}

	for i := 0; i<100; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			b1.Deposit(10)
		}()
	}

	for i := 0; i<100; i++{
		wg.Add(1)
		go func(){
			defer wg.Done()
			b1.Withdraw(5)
		}()
	}

	wg.Wait()
	fmt.Println(b1.balance)
}
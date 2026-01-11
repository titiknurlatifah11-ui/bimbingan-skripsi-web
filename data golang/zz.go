package main

import (
	"errors"
	"fmt"
)

type Account struct {
	balance float64
}

func (a *Account) Withdraw(amount float64) error {
	if amount > a.balance {
		return errors.New("insufficient balance")
	}
	a.balance -= amount
	return nil
}

func main() {
	acc := Account{balance: 500}
	err := acc.Withdraw(600)
	if err != nil {
		fmt.Println("Transaction failed:", err)
	} else {
		fmt.Println("Transaction successful, remaining balance:", acc.balance)
	}
}
package main

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

//String is a refactored method, but I'm still confused about it...
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

//Deposit means add money to the wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

//Balance show the money remaining in the wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

//Withdraw means reduce money in wallet
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

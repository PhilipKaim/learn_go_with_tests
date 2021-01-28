package main

import (
	"errors"
	"fmt"
)

// Bitcoin is an integer type to represent the balance in users Wallet
type Bitcoin int

// Wallet is a struct containing the state of the users Wallet
type Wallet struct {
	balance Bitcoin
}

// Deposit adds bitcoin to the Wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Balance checks the current amount of bitcoin in the Wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Stringer is an interface that is a contract to Bitcoin to create a
// string out of the btc value
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// ErrInsufficientFunds is an error for when funds are insufficient
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// Withdraw is a method on the Wallet struct to check if funds are sufficient
// and if they are withdraw the funds
func (w *Wallet) Withdraw(amount Bitcoin) error {

	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}

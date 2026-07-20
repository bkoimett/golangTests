package main

import (
	"fmt"
)

type Wallet struct {
	Owner string
	Balance float64
}

func (w *Wallet) AddMoney(amount float64) {
	w.Balance += amount
}

func main() {
	myWallet := Wallet{"Benjamin",50.0}

	myWallet.AddMoney(25.0)

	fmt.Println(myWallet.Balance)
}

// this is to test my repo
func test()
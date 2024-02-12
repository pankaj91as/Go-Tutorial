package main

import "fmt"

type BankingOperations interface {
	MakeDeposit(amount int)
	DoWithdraw(amount int) error
	GetBalance() int
}

func main() {
	ax := NewAxisBankDetails()
	ax.MakeDeposit(1000)

	err := ax.DoWithdraw(350)
	if err != nil {
		panic(err)
	}

	bal := ax.GetBalance()
	fmt.Printf("Balance:%d\n", bal)

	sw := NewSwissBankDetails()
	sw.MakeDeposit(1000)

	err = sw.DoWithdraw(350)
	if err != nil {
		panic(err)
	}

	bal = sw.GetBalance()
	fmt.Printf("Balance:%d\n", bal)
}

package main

import (
	"errors"
)

type AxisBankDetails struct {
	balance int
}

func NewAxisBankDetails() *AxisBankDetails {
	return &AxisBankDetails{
		balance: 0,
	}
}

func (ax *AxisBankDetails) MakeDeposit(a int) {
	ax.balance += a
}

func (ax *AxisBankDetails) DoWithdraw(a int) error {
	if err := a > ax.balance; err {
		return errors.New("you dont have enough balance")
	}

	ax.balance -= a
	return nil
}

func (ax *AxisBankDetails) GetBalance() int {
	return ax.balance
}

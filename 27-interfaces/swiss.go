package main

import (
	"errors"
)

type SwissBankDetails struct {
	balance          int
	conversionCharge int
}

func NewSwissBankDetails() *SwissBankDetails {
	return &SwissBankDetails{
		balance:          0,
		conversionCharge: 2,
	}
}

func (ax *SwissBankDetails) MakeDeposit(a int) {
	ax.balance += a
}

func (ax *SwissBankDetails) DoWithdraw(a int) error {
	if err := (a + ((a / 100) * ax.conversionCharge)) > ax.balance; err {
		return errors.New("you dont have enough balance")
	}

	ax.balance -= a + ((a / 100) * ax.conversionCharge)
	return nil
}

func (ax *SwissBankDetails) GetBalance() int {
	return ax.balance
}

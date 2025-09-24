package main

import (
	"fmt"

	"example.com/bank/bank_account"
)

func main() {
	account := bank_account.NewBackAccountDetails(0)
	account.CreditAccount(100)
	account.CreditAccount(100)
	fmt.Println(account.Balance, account.Transactions)
}

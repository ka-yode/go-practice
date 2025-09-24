package bank_account

import (
	"errors"
	"fmt"
	"strconv"
	s "strings"
	"time"
)

type transaction struct {
	amount int64
}
type BankAccountDetails struct {
	Balance      int64
	Transactions map[string]transaction
}

func NewBackAccountDetails(initialAmout int64) *BankAccountDetails {
	return &BankAccountDetails{
		Balance: initialAmout,
	}
}

func (bank *BankAccountDetails) DebitAccount(amount int64) (int64, map[string]transaction, error) {
	if amount > bank.Balance || amount <= 0 {
		return 0, bank.Transactions, errors.New("couldn't debit account")
	}
	bank.Balance = bank.Balance - amount

	if bank.Transactions == nil {
		bank.Transactions = make(map[string]transaction)
	}

	bank.Transactions[time.Now().String()] = transaction{amount: -(amount)}
	return bank.Balance, bank.Transactions, nil
}

func (bank *BankAccountDetails) CreditAccount(amount int64) (int64, map[string]transaction, error) {
	if amount <= 0 {
		return 0, bank.Transactions, errors.New("you must input a valid credit amount")
	}
	bank.Balance = bank.Balance + amount

	if bank.Transactions == nil {
		bank.Transactions = make(map[string]transaction)
	}
	bank.Transactions[time.Now().String()] = transaction{amount: amount}
	return bank.Balance, bank.Transactions, nil
}

func (bank *BankAccountDetails) BatchTransactions(transactionString string) (int64, error) {
	splitTransactions := s.Split(transactionString, ":")
	for _, value := range splitTransactions {

		isSubtract := s.HasPrefix(value, "-")
		if isSubtract {
			numberOnly := s.ReplaceAll(value, "-", "")
			number, err := convertStringToFloat(numberOnly)
			if err != nil {
				return 0, errors.New("cant complete batch transaction")

			}
			fmt.Printf("debitting account with %v\n", number)
			newBalance, _, err := bank.DebitAccount(number)
			if err != nil {
				return 0, errors.New("cant complete batch transaction")
			} else {
				fmt.Printf("new balance %v\n", newBalance)
			}

		} else {

			number, err := convertStringToFloat(value)
			if err != nil {
				return 0, errors.New("cant complete batch transaction")
			}
			fmt.Printf("crediting account with %v\n", number)
			newBalance, _, err := bank.CreditAccount(number)
			if err != nil {
				return 0, errors.New("cant complete batch transaction")
			} else {
				fmt.Printf("new balance %v\n", newBalance)
			}
		}
	}
	fmt.Println("Transactions: ", bank.Transactions)
	return bank.Balance, nil
}

func convertStringToFloat(value string) (int64, error) {
	trimmedNumberOnly := s.Trim(value, " ")
	number, err := strconv.ParseInt(trimmedNumberOnly, 0, 64)
	if err != nil {
		return 0, errors.New("cant convert string to float")
	}
	return number, nil

}

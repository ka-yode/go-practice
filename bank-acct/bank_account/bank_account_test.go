package bank_account_test

import (
	"fmt"
	"testing"

	"example.com/bank/bank_account"
)

func TestDebitAccount(t *testing.T) {
	t.Run("Shoud reduce balance [Positive number]", func(t *testing.T) {
		account := bank_account.BankAccountDetails{
			Balance: 200,
		}
		newBalance, transactions, err := account.DebitAccount(100)
		if err != nil {
			t.Fatal(err)
		}
		var expectedBalance int64 = 100

		fmt.Println(transactions)
		if newBalance != expectedBalance {
			t.Errorf("Expected %d but got %d", expectedBalance, newBalance)
		}
		fmt.Println("transactions: ", transactions)
	})
	t.Run("Shoud reduce balance [Negative number]", func(t *testing.T) {
		account := bank_account.BankAccountDetails{
			Balance: 200,
		}

		_, _, err := account.DebitAccount(-100)
		if err == nil {
			t.Fatal(err)
		}

	})
	t.Run("Shoud reduce balance [0]", func(t *testing.T) {
		account := bank_account.BankAccountDetails{
			Balance: 200,
		}

		_, _, err := account.DebitAccount(0)
		if err == nil {
			t.Fatal(err)
		}

	})

}

func TestCreditAccount(t *testing.T) {
	t.Run("Shoud increase balance [Positive number]", func(t *testing.T) {
		account := bank_account.BankAccountDetails{
			Balance: 200,
		}
		newBalance, _, err := account.CreditAccount(100)
		if err != nil {
			t.Fatal(err)
		}
		var expectedBalance int64 = 300

		if newBalance != expectedBalance {
			t.Errorf("Expected %d but got %d", expectedBalance, newBalance)
		}

	})
	t.Run("Shoud increase balance [Negative number]", func(t *testing.T) {
		account := bank_account.BankAccountDetails{
			Balance: 200,
		}

		_, _, err := account.CreditAccount(-100)
		if err == nil {
			t.Fatal(err)
		}

	})
	t.Run("Shoud increase balance [0]", func(t *testing.T) {
		account := bank_account.BankAccountDetails{
			Balance: 200,
		}

		_, _, err := account.CreditAccount(0)
		if err == nil {
			t.Fatal(err)
		}

	})

}

func TestBatch(t *testing.T) {
	t.Run("all string input", func(t *testing.T) {
		account := bank_account.BankAccountDetails{
			Balance: 200,
		}
		_, err := account.BatchTransactions("kayode")
		if err == nil {
			t.Fatal(err)
		}

	})
	t.Run("all positive input", func(t *testing.T) {
		account := bank_account.BankAccountDetails{
			Balance: 200,
		}

		newBalance, err := account.BatchTransactions("200:200:200")
		if err != nil {
			t.Fatal(err)
		}
		var expectedBalance int64 = 800
		if newBalance != expectedBalance {
			t.Errorf("Expected %d but got %d", expectedBalance, newBalance)
		}

	})
	t.Run("all negative input", func(t *testing.T) {
		account := bank_account.BankAccountDetails{
			Balance: 1000,
		}

		newBalance, err := account.BatchTransactions("-200:-200:-200")
		if err != nil {
			t.Fatal(err)
		}
		var expectedBalance int64 = 400
		if newBalance != expectedBalance {
			t.Errorf("Expected %d but got %d", expectedBalance, newBalance)
		}

	})
	t.Run("mixed input", func(t *testing.T) {
		account := bank_account.BankAccountDetails{
			Balance: 1000,
		}

		newBalance, err := account.BatchTransactions("-200:200:-200:150")
		if err != nil {
			t.Fatal(err)
		}
		var expectedBalance int64 = 950
		if newBalance != expectedBalance {
			t.Errorf("Expected %d but got %d", expectedBalance, newBalance)
		}

	})
	t.Run("mixed aield input", func(t *testing.T) {
		account := bank_account.BankAccountDetails{
			Balance: 1000,
		}

		_, err := account.BatchTransactions("-500:700:-400:-100:-500:-400")
		if err == nil {
			t.Fatal(err)
		}

	})

}

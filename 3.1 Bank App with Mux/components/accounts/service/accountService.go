package service

import (
	"errors"
)

type Account struct {
	ID       int     `json:"id"`
	BankID   int     `json:"bankId"`
	Balance  float64 `json:"balance"`
	IsActive bool    `json:"isActive"`
	Passbook []*Transaction
}

var accounts []*Account

var (
	errInsufficientBalance  = errors.New("insufficient balance")
	errInvalidDepositAmount = errors.New("deposit amount should be greater than 0")
)

func GetActivityStatus(a *Account) bool {
	return a.IsActive
}

func GetAccountByID(id int) (*Account, error) {
	for _, acc := range accounts {
		if acc.ID == id {
			return acc, nil
		}
	}
	return nil, errors.New("account not found")
}

func CreateAccount(bankID int, balance float64) (*Account, error) {
	if balance < 1000 {
		return nil, errors.New("initial balance must be at least Rs. 1000")
	}

	newAccount := &Account{
		ID:       len(accounts) + 1,
		BankID:   bankID,
		Balance:  balance,
		IsActive: true,
		Passbook: []*Transaction{},
	}

	initialTransaction := NewTransaction(0, "Initial Deposit", balance, balance, -1, -1)
	newAccount.Passbook = append(newAccount.Passbook, initialTransaction)

	accounts = append(accounts, newAccount)
	return newAccount, nil
}

// Delete an account
func DeleteAccount(a *Account) {
	a.IsActive = false
}

// Deposit adds money to the account
func Deposit(account *Account, amount float64) error {
	if amount <= 0 {
		return errInvalidDepositAmount
	}

	account.Balance += amount
	toTransaction := NewTransaction(len(account.Passbook)+1, "Credit", amount, account.Balance, -1, -1)
	account.Passbook = append(account.Passbook, toTransaction)
	return nil
}

// Withdraw removes money from the account if there are sufficient funds
func Withdraw(account *Account, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}

	if account.Balance < amount {
		return errInsufficientBalance
	}

	account.Balance -= amount
	fromTransaction := NewTransaction(len(account.Passbook)+1, "Debit", amount, account.Balance, -1, -1)
	account.Passbook = append(account.Passbook, fromTransaction)
	return nil
}

func GetAccountPassbook(a *Account) []*Transaction {
	return a.Passbook
}

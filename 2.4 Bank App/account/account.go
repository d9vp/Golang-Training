package account

import (
	"bankingApp/transactions"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type AccountFunctions interface {
	RemoveAccount()
	GetPassbook()
	Deposit(amount float64, ids ...string) error
	Withdraw(amount float64, ids ...string) error
	Transfer(toAccount AccountFunctions, amount float64) error
	GetAccountNumber() int
	GetBankID() int
	GetActivityStatus() bool
	GetBalance() float64
}

type Account struct {
	AccountNo int
	BankID    int
	Balance   float64
	IsActive  bool
	Passbook  []transactions.TransactionInterface
}

var (
	errInsufficientBalance  = errors.New("insufficient balance")
	errInvalidDepositAmount = errors.New("deposit amount should be greater than 0")
)

func NewAccount(accountNo, bankID int, firstPayment float64) *Account {
	firstTransaction := transactions.NewTransaction(0, "Credit", firstPayment, firstPayment, "self", "self")
	return &Account{
		AccountNo: accountNo,
		BankID:    bankID,
		Balance:   firstPayment,
		IsActive:  true,
		Passbook:  []transactions.TransactionInterface{firstTransaction},
	}
}

func (a *Account) GetAccountNumber() int {
	return a.AccountNo
}

func (a *Account) GetBankID() int {
	return a.BankID
}

func (a *Account) GetActivityStatus() bool {
	return a.IsActive
}

func (a *Account) GetBalance() float64 {
	return a.Balance
}

func (a *Account) Deposit(amount float64, ids ...string) error {
	if amount <= 0 {
		return errInvalidDepositAmount
	}
	a.Balance += amount
	transaction := a.createTransaction("Credit", amount, ids...)
	a.Passbook = append(a.Passbook, transaction)
	fmt.Printf("Deposited Rs. %.2f to account %d\n", amount, a.AccountNo)
	return nil
}

func (a *Account) Withdraw(amount float64, ids ...string) error {
	if amount > a.Balance {
		return errInsufficientBalance
	}
	a.Balance -= amount
	transaction := a.createTransaction("Debit", amount, ids...)
	a.Passbook = append(a.Passbook, transaction)
	fmt.Printf("Withdrew Rs. %.2f from account %d\n", amount, a.AccountNo)
	return nil
}

func (a *Account) Transfer(toAccount AccountFunctions, amount float64) error {
	if err := a.Withdraw(amount, strconv.Itoa(toAccount.GetBankID()), strconv.Itoa(toAccount.GetAccountNumber())); err != nil {
		return err
	}
	return toAccount.Deposit(amount, strconv.Itoa(a.BankID), strconv.Itoa(a.AccountNo))
}

func (a *Account) RemoveAccount() {
	a.IsActive = false
}

func (a *Account) GetPassbook() {
	fmt.Printf("%-15s %-15s %-10s %-15s %-20s %-20s %-20s\n", "Transaction ID", "Type", "Amount", "New Balance", "Correspondent Bank", "Correspondent Account", "Timestamp")
	fmt.Println(strings.Repeat("-", 110))

	for _, tran := range a.Passbook {
		tran.GetTransaction()
	}
}

func (a *Account) createTransaction(transactionType string, amount float64, ids ...string) transactions.TransactionInterface {
	tranID := 0
	if len(a.Passbook) != 0 {
		tranID = a.Passbook[len(a.Passbook)-1].GetTransactionID() + 1
	}

	if len(ids) == 0 {
		return transactions.NewTransaction(tranID, transactionType, amount, a.Balance, "self", "self")
	}

	return transactions.NewTransaction(tranID, transactionType, amount, a.Balance, ids[0], ids[1])
}

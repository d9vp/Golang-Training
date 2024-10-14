package account

import (
	"bankingApp/transactions"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Account represents a bank account.
type Account struct {
	AccountNo int
	BankID    int
	Balance   float64
	IsActive  bool
	Passbook  []*transactions.Transaction
}

// Error messages
var (
	errInsufficientBalance  = errors.New("insufficient balance")
	errInvalidDepositAmount = errors.New("deposit amount should be greater than 0")
	errInactiveAccount      = errors.New("account is inactive")
)

// NewAccount creates a new account with the given details.
func NewAccount(accountNo, bankID int, firstPayment float64) *Account {
	firstTransaction := transactions.NewTransaction(0, "Credit", firstPayment, firstPayment, "self", "self")
	return &Account{
		AccountNo: accountNo,
		BankID:    bankID,
		Balance:   firstPayment,
		IsActive:  true,
		Passbook:  []*transactions.Transaction{firstTransaction},
	}
}

// GetAccountNumber returns the account number.
func (a *Account) GetAccountNumber() int {
	return a.AccountNo
}

// GetBankID returns the bank ID associated with the account.
func (a *Account) GetBankID() int {
	return a.BankID
}

// GetActivityStatus returns the active status of the account.
func (a *Account) GetActivityStatus() bool {
	return a.IsActive
}

// GetBalance returns the current balance of the account.
func (a *Account) GetBalance() float64 {
	return a.Balance
}

// Deposit adds the specified amount to the account balance.
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

// Withdraw removes the specified amount from the account balance.
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

// Transfer moves the specified amount from this account to another account.
func (a *Account) Transfer(toAccount *Account, amount float64) error {
	if err := a.Withdraw(amount, strconv.Itoa(toAccount.BankID), strconv.Itoa(toAccount.AccountNo)); err != nil {
		return err
	}
	return toAccount.Deposit(amount, strconv.Itoa(a.BankID), strconv.Itoa(a.AccountNo))
}

// RemoveAccount deactivates the account.
func (a *Account) RemoveAccount() {
	a.IsActive = false
}

// GetPassbook displays the transaction history in a tabular format.
func (a *Account) GetPassbook() {
	fmt.Printf("%-15s %-15s %-10s %-15s %-20s %-20s %-20s\n", "Transaction ID", "Type", "Amount", "New Balance", "Correspondent Bank", "Correspondent Account", "Timestamp")
	fmt.Println(strings.Repeat("-", 110))

	for _, tran := range a.Passbook {
		fmt.Printf("%-15d %-15s %-10.2f %-15.2f %-20s %-20s %-20s\n",
			tran.TransactionID,
			tran.TransactionType,
			tran.Amount,
			tran.NewBalance,
			tran.BankIDOfCorrespondent,
			tran.AccountIDOfCorrespondent,
			tran.TimeStamp.Format("2006-01-02 15:04:05"),
		)
	}
}

// createTransaction creates a new transaction with the current account state.
func (a *Account) createTransaction(transactionType string, amount float64, ids ...string) *transactions.Transaction {
	tranID := 0
	if len(a.Passbook) != 0 {
		tranID = a.Passbook[len(a.Passbook)-1].GetTransactionID() + 1
	}

	if len(ids) == 0 {
		return transactions.NewTransaction(tranID, transactionType, amount, a.Balance, "self", "self")
	}
	return transactions.NewTransaction(tranID, transactionType, amount, a.Balance, ids[0], ids[1])
}

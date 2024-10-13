package account

import (
	"bankingApp/transactions"
	"fmt"
	"strconv"
	"strings"
)

type Account struct {
	AccountNo int
	BankID    int
	Balance   float64
	IsActive  bool
	Passbook  []*transactions.Transaction
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

func NewAccount(accountNo, bankID int, firstPayment float64) *Account {
	//validation carried out in the process of calling

	firstTransaction := transactions.NewTransaction(0, "Credit", firstPayment, firstPayment, "self", "self")
	var passbook []*transactions.Transaction
	passbook = append(passbook, firstTransaction)
	return &Account{
		AccountNo: accountNo,
		BankID:    bankID,
		Balance:   firstPayment,
		IsActive:  true,
		Passbook:  passbook,
	}
}

func (a *Account) Deposit(amount float64, ids ...string) {
	if amount <= 0 {
		fmt.Println("Deposit amount should be greater than 0")
		return
	}
	a.Balance += amount
	tranID := 0
	if len(a.Passbook) != 0 {
		tranID = a.Passbook[len(a.Passbook)-1].GetTransactionID() + 1
	}
	var tempTransaction *transactions.Transaction
	if ids == nil {
		tempTransaction = transactions.NewTransaction(tranID, "Credit", amount, a.Balance, "self", "self")
	} else {
		tempTransaction = transactions.NewTransaction(tranID, "Credit", amount, a.Balance, ids[0], ids[1])
	}

	a.Passbook = append(a.Passbook, tempTransaction)
	fmt.Printf("Deposited Rs. %.2f to account %d\n", amount, a.AccountNo)
}

func (a *Account) Withdraw(amount float64, ids ...string) {
	if amount > a.Balance {
		fmt.Println("Insufficient balance for withdrawal")
		return
	}
	a.Balance -= amount
	tranID := 0
	if len(a.Passbook) != 0 {
		tranID = a.Passbook[len(a.Passbook)-1].GetTransactionID() + 1
	}
	var tempTransaction *transactions.Transaction
	if ids == nil {
		tempTransaction = transactions.NewTransaction(tranID, "Debit", amount, a.Balance, "self", "self")
	} else {
		tempTransaction = transactions.NewTransaction(tranID, "Debit", amount, a.Balance, ids[0], ids[1])
	}
	// tempTransaction := transactions.NewTransaction(tranID, "Debit", amount, a.Balance, "self", "self")
	a.Passbook = append(a.Passbook, tempTransaction)
	fmt.Printf("Withdrew Rs. %.2f from account %d\n", amount, a.AccountNo)
}

func (a *Account) Transfer(toAccount *Account, amount float64) {
	if amount > a.Balance {
		fmt.Println("Insufficient balance for transfer")
		return
	}
	toBankID := strconv.Itoa(toAccount.BankID)
	toAccNo := strconv.Itoa(toAccount.AccountNo)
	fromBankID := strconv.Itoa(a.BankID)
	fromAccNo := strconv.Itoa(a.AccountNo)
	a.Withdraw(amount, toBankID, toAccNo)
	toAccount.Deposit(amount, fromBankID, fromAccNo)
	fmt.Printf("Transferred Rs. %.2f from account %d to account %d\n", amount, a.AccountNo, toAccount.AccountNo)
}

func (a *Account) RemoveAccount() {
	a.IsActive = false
}

func (a *Account) GetPassbook() {
	fmt.Printf("%-15s %-15s %-10s %-15s %-20s %-20s %-20s\n", "Transaction ID", "Type", "Amount", "New Balance", "Correspondent Bank", "Correspondent Account", "Timestamp")
	fmt.Println(strings.Repeat("-", 110)) // Line for header separation

	// Print each transaction in a table format
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

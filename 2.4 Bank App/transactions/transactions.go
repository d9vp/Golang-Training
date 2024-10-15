package transactions

import (
	"fmt"
	"time"
)

type TransactionInterface interface {
	GetTransactionID() int
	GetTransaction()
}

type Transaction struct {
	TransactionID            int
	TransactionType          string
	Amount                   float64
	NewBalance               float64
	BankIDOfCorrespondent    string
	AccountIDOfCorrespondent string
	TimeStamp                time.Time
}

func NewTransaction(transactionID int, transactionType string, amount, newBalance float64, bankIDOfCorr, accountIDOfCorr string) TransactionInterface {
	return &Transaction{
		TransactionID:            transactionID,
		TransactionType:          transactionType,
		Amount:                   amount,
		NewBalance:               newBalance,
		BankIDOfCorrespondent:    bankIDOfCorr,
		AccountIDOfCorrespondent: accountIDOfCorr,
		TimeStamp:                time.Now(),
	}
}

func (tran *Transaction) GetTransactionID() int {
	return tran.TransactionID
}

func (tran *Transaction) GetTransaction() {
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

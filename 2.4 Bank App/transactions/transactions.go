package transactions

import (
	"time"
)

type Transaction struct {
	TransactionID            int
	TransactionType          string
	Amount                   float64
	NewBalance               float64
	BankIDOfCorrespondent    string
	AccountIDOfCorrespondent string
	TimeStamp                time.Time
}

func NewTransaction(transactionID int, transactionType string, amount, newBalance float64, bankIDofcorr, accountIDofcorr string) *Transaction {
	return &Transaction{
		TransactionID:            transactionID,
		TransactionType:          transactionType,
		Amount:                   amount,
		NewBalance:               newBalance,
		BankIDOfCorrespondent:    bankIDofcorr,
		AccountIDOfCorrespondent: accountIDofcorr,
		TimeStamp:                time.Now(),
	}
}

func (tran *Transaction) GetTransactionID() int {
	return tran.TransactionID
}

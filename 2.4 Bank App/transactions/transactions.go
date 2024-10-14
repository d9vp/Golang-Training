package transactions

import (
	"time"
)

// Transaction represents a financial transaction.
type Transaction struct {
	TransactionID            int
	TransactionType          string
	Amount                   float64
	NewBalance               float64
	BankIDOfCorrespondent    string
	AccountIDOfCorrespondent string
	TimeStamp                time.Time
}

// NewTransaction creates a new transaction with the given details.
func NewTransaction(transactionID int, transactionType string, amount, newBalance float64, bankIDOfCorr, accountIDOfCorr string) *Transaction {
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

// GetTransactionID returns the transaction ID.
func (tran *Transaction) GetTransactionID() int {
	return tran.TransactionID
}

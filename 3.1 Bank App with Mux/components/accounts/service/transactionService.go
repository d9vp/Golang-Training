package service

import (
	"strconv"
	"time"
)

// // Transaction struct defines the structure of a transaction
type Transaction struct {
	ID                      int       `json:"id"`
	Type                    string    `json:"type"` // e.g., "deposit", "withdrawal", "transfer"
	Amount                  float64   `json:"amount"`
	BalanceAfterTransaction float64   `json:"balanceAfterTransaction"`
	CorrespondingBankID     string    `json:"correspondingBankId"`
	CorrespondingAccountID  string    `json:"correspondingAccountId"`
	Timestamp               time.Time `json:"timestamp"`
}

// // In-memory slice to store all transactions
// var transactions []*Transaction

// // GetAllTransactions returns all transactions
// func GetAllTransactions() []*Transaction {
// 	return transactions
// }

// // GetTransactionByID returns a specific transaction by its ID
// func GetTransactionByID(id int) (*Transaction, error) {
// 	for _, tran := range transactions {
// 		if tran.ID == id {
// 			return tran, nil
// 		}
// 	}
// 	return nil, errors.New("transaction not found")
// }

// // CreateTransaction adds a new transaction to the list
// func CreateTransaction(transaction *Transaction) error {
// 	if err := validateTransaction(transaction); err != nil {
// 		return err
// 	}

// 	transaction.ID = len(transactions) + 1
// 	transaction.Timestamp = time.Now()

// 	transactions = append(transactions, transaction)
// 	return nil
// }

// // NewTransaction creates a new Transaction object with the provided details
func NewTransaction(id int, tranType string, amount, newBalance float64, bankIDOfCorr, accountIDOfCorr int) *Transaction {
	corrBankID := ""
	corrAccID := ""
	if bankIDOfCorr == -1 {
		corrBankID = "self"
		corrAccID = "self"
	} else {
		corrBankID = strconv.Itoa(bankIDOfCorr)
		corrAccID = strconv.Itoa(accountIDOfCorr)
	}
	return &Transaction{
		ID:                      id,
		Type:                    tranType,
		Amount:                  amount,
		BalanceAfterTransaction: newBalance,
		CorrespondingBankID:     corrBankID,
		CorrespondingAccountID:  corrAccID,
		Timestamp:               time.Now(),
	}
}

// ValidateTransaction ensures the transaction meets certain conditions
// func validateTransaction(transaction *Transaction) error {
// 	if transaction.Amount <= 0 {
// 		return errors.New("amount must be greater than zero")
// 	}

// 	if transaction.Type == "" {
// 		return errors.New("transaction type is required")
// 	}

// 	if transaction.CorrespondingBankID == "" {
// 		return errors.New("valid corresponding bank ID is required")
// 	}

// 	if transaction.CorrespondingAccountID == "" {
// 		return errors.New("valid corresponding account ID is required")
// 	}

// 	return nil
// }

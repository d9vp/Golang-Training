package service

import (
	"time"
	"user/models"

	"gorm.io/gorm"
)

// NewTransactionEntry creates a new TransactionEntry instance.
func NewTransactionEntry(tranType string, amount, newBalance float64, bankIDOfCorr, accountIDOfCorr int) *models.TransactionEntry {
	return &models.TransactionEntry{
		Type:                    tranType,
		Amount:                  amount,
		BalanceAfterTransaction: newBalance,
		CorrespondingBankID:     bankIDOfCorr,
		CorrespondingAccountID:  accountIDOfCorr,
		Timestamp:               time.Now(),
	}
}

// CreateTransaction saves a new transaction to the database.
func CreateTransaction(db *gorm.DB, transaction *models.TransactionEntry) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(transaction).Error; err != nil {
			return err
		}
		return nil
	})
}

// GetTransaction retrieves a transaction by its ID.
func GetTransaction(db *gorm.DB, id int) (*models.TransactionEntry, error) {
	var transaction models.TransactionEntry
	err := db.First(&transaction, id).Error
	if err != nil {
		return nil, err
	}
	return &transaction, nil
}

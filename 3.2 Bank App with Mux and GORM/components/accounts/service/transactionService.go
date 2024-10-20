package service

import (
	"time"
	"user/models"

	"gorm.io/gorm"
)

// NewTransaction creates a new Transaction instance.
func NewTransaction(tranType string, amount, newBalance float64, bankIDOfCorr, accountIDOfCorr int) *models.Transaction {
	return &models.Transaction{
		Type:                    tranType,
		Amount:                  amount,
		BalanceAfterTransaction: newBalance,
		CorrespondingBankID:     bankIDOfCorr,
		CorrespondingAccountID:  accountIDOfCorr,
		Timestamp:               time.Now(),
	}
}

// CreateTransaction saves a transaction to the database.
func CreateTransaction(db *gorm.DB, transaction *models.Transaction) error {
	return db.Create(transaction).Error
}

// GetTransaction retrieves a transaction by ID.
func GetTransaction(db *gorm.DB, id int) (*models.Transaction, error) {
	var transaction models.Transaction
	err := db.First(&transaction, id).Error
	return &transaction, err
}

// Other CRUD functions can be added here as needed

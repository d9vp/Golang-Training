package service

import (
	"errors"
	"user/models"

	"gorm.io/gorm"
)

// AddToLedger updates the ledger entries for two banks involved in a transaction using a transaction
func AddToLedger(bankID int, corrBankID int, amount float64) error {
	// Start the transaction
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		// Get the first bank within the transaction
		_, err := GetBankByID(bankID)
		if err != nil {
			return errBankNotFound
		}

		// Get the corresponding bank
		_, err = GetBankByID(corrBankID)
		if err != nil {
			return errBankNotFound
		}

		// Update ledger for the first bank
		if err := updateOrCreateLedgerEntry(tx, bankID, corrBankID, amount); err != nil {
			return errors.New("failed to update ledger for bankID")
		}

		// Update ledger for the corresponding bank
		if err := updateOrCreateLedgerEntry(tx, corrBankID, bankID, -amount); err != nil {
			return errors.New("failed to update ledger for corresponding bank")
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// updateOrCreateLedgerEntry checks if a ledger entry exists; if so, updates it; otherwise, creates a new one.
// This function is now wrapped in the same transaction.
func updateOrCreateLedgerEntry(tx *gorm.DB, bankID int, corrBankID int, amount float64) error {
	var ledger models.LedgerData

	// Check if a ledger entry exists for the given bank and corresponding bank
	err := tx.Where("bank_id = ? AND corresponding_bank_id = ?", bankID, corrBankID).First(&ledger).Error

	if err == nil {
		// Ledger entry exists; update the amount
		ledger.Amount += amount
		if err := tx.Save(&ledger).Error; err != nil {
			return errors.New("failed to update ledger entry")
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// Ledger entry does not exist; create a new one
		newLedger := &models.LedgerData{
			BankID:              bankID,
			CorrespondingBankID: corrBankID,
			Amount:              amount,
		}
		if err := tx.Create(newLedger).Error; err != nil {
			return errors.New("failed to create new ledger entry")
		}
	} else {
		// If there's another error, return it
		return err
	}

	return nil
}

// GetLedger returns the ledger data for a particular bank using a transaction
func GetLedger(bankID int) ([]*models.LedgerData, error) {
	var ledgerData []*models.LedgerData

	// Start the transaction to retrieve ledger data
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("bank_id = ?", bankID).Find(&ledgerData).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errBankNotFound
			}
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return ledgerData, nil
}

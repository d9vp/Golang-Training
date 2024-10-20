package service

import (
	"errors"
	"user/models"

	"gorm.io/gorm"
)

// AddToLedger updates the ledger entries for two banks involved in a transaction
func AddToLedger(bankID int, corrBankID int, amount float64) error {
	// Get both banks from the database
	_, err := GetBankByID(bankID)
	if err != nil {
		return errBankNotFound
	}
	_, err = GetBankByID(corrBankID)
	if err != nil {
		return errBankNotFound
	}

	// Update ledger for the first bank
	if err := updateOrCreateLedgerEntry(bankID, corrBankID, amount); err != nil {
		return err
	}

	// Update ledger for the corresponding bank
	if err := updateOrCreateLedgerEntry(corrBankID, bankID, -amount); err != nil {
		return err
	}

	return nil
}

// updateOrCreateLedgerEntry checks if a ledger entry exists; if so, updates it; otherwise, creates a new one.
func updateOrCreateLedgerEntry(bankID int, corrBankID int, amount float64) error {
	// Check if a ledger entry exists for the given bank and corresponding bank
	var ledger models.LedgerData
	err := models.DB.Where("bank_id = ? AND corresponding_bank_id = ?", bankID, corrBankID).First(&ledger).Error

	if err == nil {
		// Ledger entry exists; update the amount
		ledger.Amount += amount
		if err := models.DB.Save(&ledger).Error; err != nil {
			return err
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		// Ledger entry does not exist; create a new one
		newLedger := &models.LedgerData{
			BankID:              bankID,
			CorrespondingBankID: corrBankID,
			Amount:              amount,
		}
		if err := models.DB.Create(newLedger).Error; err != nil {
			return err
		}
	} else {
		// If there's another error, return it
		return err
	}

	return nil
}

// GetLedger returns the ledger data for a particular bank
func GetLedger(bankID int) ([]*models.LedgerData, error) {
	var ledgerData []*models.LedgerData

	if err := models.DB.Where("bank_id = ?", bankID).Find(&ledgerData).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errBankNotFound
		}
		return nil, err
	}

	return ledgerData, nil
}

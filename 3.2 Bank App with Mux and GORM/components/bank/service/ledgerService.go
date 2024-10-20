package service

import (
	"errors"
	"user/models"

	"gorm.io/gorm"
)

func AddToLedger(bankID int, corrBankID int, amount float64) error {
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		_, err := GetBankByID(bankID)
		if err != nil {
			return errBankNotFound
		}

		_, err = GetBankByID(corrBankID)
		if err != nil {
			return errBankNotFound
		}

		if err := updateOrCreateLedgerEntry(tx, bankID, corrBankID, amount); err != nil {
			return errors.New("failed to update ledger for bankID")
		}

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

func updateOrCreateLedgerEntry(tx *gorm.DB, bankID int, corrBankID int, amount float64) error {
	var ledger models.LedgerData

	err := tx.Where("bank_id = ? AND corresponding_bank_id = ?", bankID, corrBankID).First(&ledger).Error

	if err == nil {
		ledger.Amount += amount
		if err := tx.Save(&ledger).Error; err != nil {
			return errors.New("failed to update ledger entry")
		}
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		newLedger := &models.LedgerData{
			BankID:              bankID,
			CorrespondingBankID: corrBankID,
			Amount:              amount,
		}
		if err := tx.Create(newLedger).Error; err != nil {
			return errors.New("failed to create new ledger entry")
		}
	} else {
		return err
	}

	return nil
}

func GetLedger(bankID int) ([]*models.LedgerData, error) {
	var ledgerData []*models.LedgerData

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

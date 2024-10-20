package service

import (
	"errors"
	"user/models"

	"gorm.io/gorm"
)

var (
	errBankNotFound     = errors.New("bank not found")
	errInvalidBankName  = errors.New("invalid bank name or abbreviation")
	errBankInactive     = errors.New("bank is inactive")
	errInvalidParameter = errors.New("invalid parameter for update")
)

func GetAllBanks() ([]*models.Bank, error) {
	var banks []*models.Bank
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Find(&banks).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return banks, nil
}

func GetBankByID(id int) (*models.Bank, error) {
	var bank models.Bank
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&bank, id).Error; err != nil {
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
	return &bank, nil
}

func CreateBank(bank *models.Bank) (*models.Bank, error) {
	if bank.FullName == "" || bank.Abbreviation == "" {
		return nil, errInvalidBankName
	}

	err := models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&bank).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return bank, nil
}

func UpdateBank(id int, parameter string, newValue string) error {
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		var bank models.Bank
		if err := tx.First(&bank, id).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errBankNotFound
			}
			return err
		}

		switch parameter {
		case "Full Name":
			if newValue == "" {
				return errors.New("full name cannot be empty")
			}
			bank.FullName = newValue
		case "Abbreviation":
			if len(newValue) == 0 {
				return errors.New("abbreviation cannot be empty")
			} else if len(newValue) > 5 {
				return errors.New("abbreviation too long, must be under 5 characters")
			}
			bank.Abbreviation = newValue
		default:
			return errInvalidParameter
		}

		if err := tx.Save(&bank).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func RemoveBank(bankID int) error {
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		var bank models.Bank
		if err := tx.First(&bank, bankID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errBankNotFound
			}
			return err
		}

		if !bank.IsActive {
			return errBankInactive
		}

		bank.IsActive = false

		if err := tx.Save(&bank).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func AddAccountToBank(bankID int, account *models.Account) error {
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		var bank models.Bank
		if err := tx.First(&bank, bankID).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errBankNotFound
			}
			return err
		}

		if !bank.IsActive {
			return errBankInactive
		}

		bank.Accounts = append(bank.Accounts, account)

		if err := tx.Save(bank).Error; err != nil {
			return errors.New("failed to update bank with new account")
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

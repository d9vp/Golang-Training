package service

import (
	"errors"
	"user/components/accounts/service"
	"user/models"

	"gorm.io/gorm"
)

// Error messages
var (
	errBankNotFound     = errors.New("bank not found")
	errInvalidBankName  = errors.New("invalid bank name or abbreviation")
	errBankInactive     = errors.New("bank is inactive")
	errInvalidParameter = errors.New("invalid parameter for update")
)

// Get all banks from the database
func GetAllBanks() ([]*models.Bank, error) {
	var banks []*models.Bank
	if err := models.DB.Find(&banks).Error; err != nil {
		return nil, err
	}
	return banks, nil
}

// Get a specific bank by ID
func GetBankByID(id int) (*models.Bank, error) {
	var bank models.Bank
	if err := models.DB.First(&bank, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errBankNotFound
		}
		return nil, err
	}
	return &bank, nil
}

// Create a new bank
func CreateBank(bank *models.Bank) (*models.Bank, error) {
	if bank.FullName == "" || bank.Abbreviation == "" {
		return nil, errInvalidBankName
	}

	if err := models.DB.Create(&bank).Error; err != nil {
		return nil, err
	}

	return bank, nil
}

// Update bank information
func UpdateBank(id int, parameter string, newValue string) error {
	var bank models.Bank
	if err := models.DB.First(&bank, id).Error; err != nil {
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

	if err := models.DB.Save(&bank).Error; err != nil {
		return err
	}
	return nil
}

// Remove (soft-delete) a bank
func RemoveBank(bankID int) error {
	var bank models.Bank
	if err := models.DB.First(&bank, bankID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errBankNotFound
		}
		return err
	}

	if !bank.IsActive {
		return errBankInactive
	}

	bank.IsActive = false
	// Soft-delete associated accounts
	for _, account := range bank.Accounts {
		service.DeleteAccount(account) // Call to account service to remove accounts
	}

	if err := models.DB.Save(&bank).Error; err != nil {
		return err
	}
	return nil
}

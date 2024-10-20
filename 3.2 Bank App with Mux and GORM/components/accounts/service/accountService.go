package service

import (
	"errors"
	"user/models"

	"gorm.io/gorm"
)

var (
	errInsufficientBalance  = errors.New("insufficient balance")
	errInvalidDepositAmount = errors.New("deposit amount should be greater than 0")
	errAccountNotFound      = errors.New("account not found")
	errInitialBalance       = errors.New("initial balance must be at least Rs. 1000")
)

// GetActivityStatus checks if the account is active
func GetActivityStatus(a *models.Account) bool {
	return a.IsActive
}

// GetAccountByID retrieves an account from the database by its ID
func GetAccountByID(id int) (*models.Account, error) {
	var account models.Account
	if err := models.DB.First(&account, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errAccountNotFound
		}
		return nil, err
	}
	return &account, nil
}

// CreateAccount creates a new account with an initial balance and saves it to the database
func CreateAccount(userID, bankID int, balance float64) (*models.Account, error) {
	if balance < 1000 {
		return nil, errInitialBalance
	}

	newAccount := &models.Account{
		UserID:   userID,
		BankID:   bankID,
		Balance:  balance,
		IsActive: true,
	}

	// Save account to the database
	if err := models.DB.Create(newAccount).Error; err != nil {
		return nil, err
	}

	// Create initial transaction and save it to the passbook
	initialTransaction := NewTransaction("Initial Deposit", balance, balance, -1, -1)
	if err := models.DB.Model(newAccount).Association("Passbook").Append(initialTransaction); err != nil {
		return nil, err
	}

	return newAccount, nil
}

// DeleteAccount deactivates the account by setting its status to inactive
func DeleteAccount(a *models.Account) error {
	a.IsActive = false
	return models.DB.Save(a).Error
}

// Deposit adds money to the account and updates the passbook
func Deposit(account *models.Account, amount float64) error {
	if amount <= 0 {
		return errInvalidDepositAmount
	}

	account.Balance += amount
	transaction := NewTransaction("Deposit", amount, amount, -1, -1)

	// Save changes and transaction
	if err := models.DB.Save(account).Error; err != nil {
		return err
	}
	if err := models.DB.Model(account).Association("Passbook").Append(transaction); err != nil {
		return err
	}

	return nil
}

// Withdraw removes money from the account if there are sufficient funds and updates the passbook
func Withdraw(account *models.Account, amount float64) error {
	if amount <= 0 {
		return errors.New("withdrawal amount must be greater than zero")
	}

	if account.Balance < amount {
		return errInsufficientBalance
	}

	account.Balance -= amount
	transaction := NewTransaction("Withdraw", amount, amount, -1, -1)

	// Save changes and transaction
	if err := models.DB.Save(account).Error; err != nil {
		return err
	}
	if err := models.DB.Model(account).Association("Passbook").Append(transaction); err != nil {
		return err
	}

	return nil
}

// GetAccountPassbook retrieves the transaction history (passbook) for a specific account
func GetAccountPassbook(a *models.Account) ([]models.Transaction, error) {
	var passbook []models.Transaction
	if err := models.DB.Model(a).Association("Passbook").Find(&passbook); err != nil {
		return nil, err
	}
	return passbook, nil
}

package service

import (
	"errors"
	"user/components/bank/service"
	"user/models"

	"gorm.io/gorm"
)

var (
	errInsufficientBalance = errors.New("insufficient balance")
	errAccountNotFound     = errors.New("account not found")
	errAmountLessThanZero  = errors.New("amount has to be greater than zero")
	errUserNotFound        = errors.New("user not found")
)

// GetActivityStatus checks if the account is active
func GetActivityStatus(a *models.Account) bool {
	return a.IsActive
}
func CreateAccount(user *models.User, bankID int, balance float64) (*models.Account, error) {
	var newAccount *models.Account

	err := models.DB.Transaction(func(tx *gorm.DB) error {
		// Create new account
		newAccount = &models.Account{
			UserID:   user.UserID,
			BankID:   bankID,
			Balance:  balance,
			IsActive: true,
		}

		// Save the new account to the database
		if err := tx.Create(newAccount).Error; err != nil {
			return err
		}

		// Create initial transaction and save it to the passbook
		initialTransaction := NewTransactionEntry("Initial Deposit", balance, balance, -1, -1)

		// Use Association to save passbook
		if err := tx.Model(newAccount).Association("Passbook").Append(initialTransaction); err != nil {
			return err
		}

		// Associate the account with the user
		user.Accounts = append(user.Accounts, newAccount)

		// Update the user to persist the new account association
		if err := tx.Save(user).Error; err != nil {
			return errors.New("failed to update user with new account")
		}

		// Add the account to the bank
		if err := service.AddAccountToBank(bankID, newAccount); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return newAccount, nil
}

// GetAccountsForUser retrieves all accounts for a given user
func GetAccountsForUser(user *models.User) ([]*models.Account, error) {
	var accounts []*models.Account

	if err := models.DB.Where("user_id = ?", user.UserID).Find(&accounts).Error; err != nil {
		return nil, err
	}

	return accounts, nil
}

// GetTotalBalanceForUser calculates the total balance across all accounts for a given user
func GetTotalBalanceForUser(userName string) (float64, error) {
	user, err := GetUserWithAccounts(userName)
	if err != nil {
		return 0, err
	}

	totalBalance := 0.0
	for _, acc := range user.Accounts {
		totalBalance += acc.Balance
	}

	return totalBalance, nil
}

// GetUserWithAccounts retrieves a user along with their accounts
func GetUserWithAccounts(userName string) (*models.User, error) {
	var user models.User

	if err := models.DB.Preload("Accounts").Where("user_name = ?", userName).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// GetAccountByID retrieves an account by its ID and ensures it belongs to the user
func GetAccountByID(user *models.User, accountID int) (*models.Account, error) {
	for _, account := range user.Accounts {
		if account.ID == accountID {
			if err := models.DB.Preload("Passbook").First(&account).Error; err != nil {
				return nil, err
			}
			return account, nil
		}
	}
	return nil, errAccountNotFound
}

// DeleteAccount deactivates an account by setting its status to inactive
func DeleteAccount(a *models.Account) error {
	return models.DB.Transaction(func(tx *gorm.DB) error {
		a.IsActive = false
		return tx.Save(a).Error
	})
}

// Deposit adds money to an account and updates the passbook
func Deposit(account *models.Account, amount float64) error {
	if amount <= 0 {
		return errAmountLessThanZero
	}

	return models.DB.Transaction(func(tx *gorm.DB) error {
		// Update balance
		account.Balance += amount
		transaction := NewTransactionEntry("Deposit", amount, account.Balance, -1, -1)
		account.Passbook = append(account.Passbook, transaction)

		// Save changes to account and transaction history
		if err := tx.Save(account).Error; err != nil {
			return err
		}
		if err := tx.Model(account).Association("Passbook").Append(transaction); err != nil {
			return err
		}

		return nil
	})
}

// Withdraw removes money from an account if there are sufficient funds and updates the passbook
func Withdraw(account *models.Account, amount float64) error {
	if amount <= 0 {
		return errAmountLessThanZero
	}

	if account.Balance < amount {
		return errInsufficientBalance
	}

	return models.DB.Transaction(func(tx *gorm.DB) error {
		// Update balance
		account.Balance -= amount
		transaction := NewTransactionEntry("Withdraw", amount, account.Balance, -1, -1)
		account.Passbook = append(account.Passbook, transaction)

		// Save changes to account and transaction history
		if err := tx.Save(account).Error; err != nil {
			return err
		}
		if err := tx.Model(account).Association("Passbook").Append(transaction); err != nil {
			return err
		}

		return nil
	})
}

// Transfer performs a transfer between two accounts, ensuring both operations are atomic
func Transfer(fromAccount, toAccount *models.Account, amount float64) error {
	if amount <= 0 {
		return errAmountLessThanZero
	}

	return models.DB.Transaction(func(tx *gorm.DB) error {
		// Withdraw from the sender's account
		if err := Withdraw(fromAccount, amount); err != nil {
			return err
		}

		// Deposit into the recipient's account
		if err := Deposit(toAccount, amount); err != nil {
			return err
		}

		return nil
	})
}

// GetAccountByIDForTransfer retrieves an account for transfer, ensuring it belongs to the user and is valid
func GetAccountByIDForTransfer(user *models.User, bankID, accountID int) (*models.Account, error) {
	for _, acc := range user.Accounts {
		if acc.ID == accountID && acc.BankID == bankID {
			if err := models.DB.Preload("Passbook").First(&acc).Error; err != nil {
				return nil, err
			}
			return acc, nil
		}
	}
	return nil, errAccountNotFound
}

func GetAccountPassbook(accountID int) ([]*models.TransactionEntry, error) {
	var passbook []*models.TransactionEntry

	err := models.DB.Transaction(func(tx *gorm.DB) error {
		// Retrieve the passbook entries associated with the account
		if err := tx.Where("account_id = ?", accountID).Find(&passbook).Error; err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return passbook, nil
}

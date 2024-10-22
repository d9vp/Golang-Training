package service

import (
	"errors"
	"user/models"

	"gorm.io/gorm"
)

var (
	errInsufficientBalance = errors.New("insufficient balance")
	errAccountNotFound     = errors.New("account not found")
	errAmountLessThanZero  = errors.New("amount has to be greater than zero")
	errUserNotFound        = errors.New("user not found")
)

func GetActivityStatus(a *models.Account) bool {
	return a.IsActive
}
func CreateAccount(user *models.User, bankID int, balance float64) (*models.Account, error) {
	var newAccount *models.Account

	err := models.DB.Transaction(func(tx *gorm.DB) error {
		newAccount = &models.Account{
			UserID:   int(user.ID),
			BankID:   bankID,
			Balance:  balance,
			IsActive: true,
		}

		if err := tx.Create(newAccount).Error; err != nil {
			return err
		}

		initialTransaction := NewTransactionEntry("Initial Deposit", balance, balance, -1, -1)

		if err := tx.Model(newAccount).Association("Passbook").Append(initialTransaction); err != nil {
			return err
		}

		user.Accounts = append(user.Accounts, newAccount)

		if err := tx.Save(user).Error; err != nil {
			return errors.New("failed to update user with new account")
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return newAccount, nil
}

func GetAccountsForUser(user *models.User) ([]*models.Account, error) {
	var accounts []*models.Account

	if err := models.DB.Where("user_id = ?", user.ID).Find(&accounts).Error; err != nil {
		return nil, err
	}

	return accounts, nil
}

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

func GetAccountByID(user *models.User, accountID int) (*models.Account, error) {
	for _, account := range user.Accounts {
		if int(account.ID) == accountID {
			if err := models.DB.Preload("Passbook").First(&account).Error; err != nil {
				return nil, err
			}
			return account, nil
		}
	}
	return nil, errAccountNotFound
}

func DeleteAccount(a *models.Account) error {
	return models.DB.Transaction(func(tx *gorm.DB) error {
		a.IsActive = false
		return tx.Save(a).Error
	})
}

func Deposit(account *models.Account, amount float64) error {
	if amount <= 0 {
		return errAmountLessThanZero
	}

	return models.DB.Transaction(func(tx *gorm.DB) error {
		account.Balance += amount
		transaction := NewTransactionEntry("Deposit", amount, account.Balance, -1, -1)
		account.Passbook = append(account.Passbook, transaction)

		if err := tx.Save(account).Error; err != nil {
			return err
		}
		if err := tx.Model(account).Association("Passbook").Append(transaction); err != nil {
			return err
		}

		return nil
	})
}

func Withdraw(account *models.Account, amount float64) error {
	if amount <= 0 {
		return errAmountLessThanZero
	}

	if account.Balance < amount {
		return errInsufficientBalance
	}

	return models.DB.Transaction(func(tx *gorm.DB) error {
		account.Balance -= amount
		transaction := NewTransactionEntry("Withdraw", amount, account.Balance, -1, -1)
		account.Passbook = append(account.Passbook, transaction)

		if err := tx.Save(account).Error; err != nil {
			return err
		}
		if err := tx.Model(account).Association("Passbook").Append(transaction); err != nil {
			return err
		}

		return nil
	})
}

func Transfer(fromAccount, toAccount *models.Account, amount float64) error {
	if amount <= 0 {
		return errAmountLessThanZero
	}

	return models.DB.Transaction(func(tx *gorm.DB) error {
		if err := Withdraw(fromAccount, amount); err != nil {
			return err
		}

		if err := Deposit(toAccount, amount); err != nil {
			return err
		}

		return nil
	})
}

func GetAccountByIDForTransfer(user *models.User, bankID, accountID int) (*models.Account, error) {
	for _, acc := range user.Accounts {
		if int(acc.ID) == accountID && acc.BankID == bankID {
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

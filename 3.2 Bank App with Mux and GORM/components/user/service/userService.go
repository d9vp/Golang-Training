package service

import (
	"errors"
	"log"
	"user/models"

	"gorm.io/gorm"
)

var (
	errNoUserFound  = errors.New("no such user found")
	errEmptyString  = errors.New("empty string cannot be used as a parameter")
	errUserExists   = errors.New("user name already in use")
	errInvalidParam = errors.New("invalid parameter entered")
)

// Find a user by username and preload accounts
func findUserByUserName(userName string) (*models.User, error) {
	var user models.User
	// Use a transaction
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		// Preload Accounts when finding the user
		if err := tx.Preload("Accounts").First(&user, "user_name = ?", userName).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return errNoUserFound
			}
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// Get all active users and preload their accounts
func GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		// Preload Accounts for all active users
		if err := tx.Where("is_active = ?", true).Preload("Accounts").Find(&users).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return users, nil
}

// Create a new Admin user
func NewAdmin(firstName, lastName, userName, password string) (*models.User, error) {
	// Check if user exists already
	existingUser, _ := findUserByUserName(userName)
	if existingUser != nil {
		return nil, errUserExists
	}

	newUser := &models.User{
		UserID:    0, // Auto-incremented by GORM
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
		IsAdmin:   true,
		IsActive:  true,
		Password:  password,
		Accounts:  []*models.Account{},
	}

	err := models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(newUser).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	log.Printf("Admin created with ID %d", newUser.UserID)
	return newUser, nil
}

// Create a new regular user
func NewUser(firstName, lastName, userName, password string) (*models.User, error) {
	// Check if user exists already
	existingUser, _ := findUserByUserName(userName)
	if existingUser != nil {
		return nil, errUserExists
	}

	newUser := &models.User{
		UserID:    0, // Auto-incremented by GORM
		FirstName: firstName,
		LastName:  lastName,
		UserName:  userName,
		IsAdmin:   false,
		IsActive:  true,
		Password:  password,
		Accounts:  []*models.Account{},
	}

	err := models.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(newUser).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	log.Printf("User created with ID %d", newUser.UserID)
	return newUser, nil
}

// Get user accounts and preload their transaction history (Passbook)
func GetAccountsForUser(u *models.User) ([]*models.Account, error) {
	var accounts []*models.Account
	err := models.DB.Transaction(func(tx *gorm.DB) error {
		// Preload transaction history for each account (Passbook)
		if err := tx.Model(u).Preload("Accounts.Passbook").Association("Accounts").Find(&accounts); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return accounts, nil
}

// Update user details
func UpdateUsers(userName, parameter string, newValue interface{}) error {
	user, err := findUserByUserName(userName)
	if err != nil {
		return err
	}

	// Update the user in a transaction
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		switch parameter {
		case "First Name":
			if value, ok := newValue.(string); ok && value != "" {
				user.FirstName = value
			} else {
				return errEmptyString
			}
		case "Last Name":
			if value, ok := newValue.(string); ok && value != "" {
				user.LastName = value
			} else {
				return errEmptyString
			}
		case "Password":
			if value, ok := newValue.(string); ok && value != "" {
				user.Password = value
			} else {
				return errEmptyString
			}
		case "Admin Rights":
			if value, ok := newValue.(bool); ok {
				user.IsAdmin = value
			}
		default:
			return errInvalidParam
		}

		// Save changes
		if err := tx.Save(&user).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// Soft-delete a user (set IsActive to false)
func DeleteUsers(userName string) error {
	user, err := findUserByUserName(userName)
	if err != nil {
		return err
	}

	// Soft delete using a transaction
	err = models.DB.Transaction(func(tx *gorm.DB) error {
		user.IsActive = false
		if err := tx.Save(&user).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

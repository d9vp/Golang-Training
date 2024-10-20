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
	errUserExists   = errors.New("user already exists")
	errInvalidParam = errors.New("invalid parameter entered")
)

// Find a user by ID
func findUserByID(userID int) (*models.User, error) {
	var user models.User
	if err := models.DB.First(&user, "user_id = ?", userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errNoUserFound
		}
		return nil, err
	}
	return &user, nil
}

// Get all users
func GetAllUsers() ([]*models.User, error) {
	var users []*models.User
	if err := models.DB.Where("is_active = ?", true).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Create a new Admin user
func NewAdmin(firstName, lastName, password string) (*models.User, error) {
	if firstName == "" || lastName == "" || password == "" {
		return nil, errEmptyString
	}

	// Check if user exists already
	existingUser, _ := findUserByName(firstName, lastName)
	if existingUser != nil {
		return nil, errUserExists
	}

	newUser := &models.User{
		UserID:    0, // Auto-incremented by GORM
		FirstName: firstName,
		LastName:  lastName,
		IsAdmin:   true,
		IsActive:  true,
		Password:  password,
		Accounts:  []*models.Account{},
	}

	if err := models.DB.Create(newUser).Error; err != nil {
		return nil, err
	}
	log.Printf("Admin created with ID %d", newUser.UserID)
	return newUser, nil
}

// Create a new regular user
func NewUser(firstName, lastName, password string) (*models.User, error) {
	if firstName == "" || lastName == "" || password == "" {
		return nil, errEmptyString
	}

	// Check if user exists already
	existingUser, _ := findUserByName(firstName, lastName)
	if existingUser != nil {
		return nil, errUserExists
	}

	newUser := &models.User{
		UserID:    0, // Auto-incremented by GORM
		FirstName: firstName,
		LastName:  lastName,
		IsAdmin:   false,
		IsActive:  true,
		Password:  password,
		Accounts:  []*models.Account{},
	}

	if err := models.DB.Create(newUser).Error; err != nil {
		return nil, err
	}
	log.Printf("User created with ID %d", newUser.UserID)
	return newUser, nil
}

// Get all active users
func GetUsers() ([]*models.User, error) {
	var users []*models.User
	if err := models.DB.Where("is_active = ?", true).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetAccountsForUser(u *models.User) []*models.Account {
	return u.Accounts
}

// Update user details
func UpdateUsers(userID int, parameter string, newValue interface{}) error {
	user, err := findUserByID(userID)
	if err != nil {
		return err
	}

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
	if err := models.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// Soft-delete a user (set IsActive to false)
func DeleteUsers(userID int) error {
	user, err := findUserByID(userID)
	if err != nil {
		return err
	}

	user.IsActive = false
	if err := models.DB.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

// Helper function to find user by name
func findUserByName(firstName, lastName string) (*models.User, error) {
	var user models.User
	if err := models.DB.First(&user, "first_name = ? AND last_name = ?", firstName, lastName).Error; err != nil {
		return nil, errNoUserFound
	}
	return &user, nil
}

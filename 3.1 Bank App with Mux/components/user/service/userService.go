package service

import (
	"errors"
	accService "user/components/accounts/service"
)

type User struct {
	UserID    int
	FirstName string
	LastName  string
	IsAdmin   bool
	IsActive  bool
	Password  string
	Accounts  []*accService.Account
}

var (
	errNoUserFound = errors.New("no such user found")
	errEmptyString = errors.New("empty string cannot be used as a parameter")
)

var superAdmin = &User{
	UserID:    0,
	FirstName: "Super",
	LastName:  "Admin",
	IsAdmin:   true,
	IsActive:  true,
	Password:  "password",
	Accounts:  nil,
}
var AllUsers = []*User{superAdmin}

func findUserID() int {
	if len(AllUsers) == 0 {
		return 0
	}
	return AllUsers[len(AllUsers)-1].UserID + 1
}

func findUserByID(userID int) *User {
	for _, cust := range AllUsers {
		if cust.UserID == userID && cust.IsActive {
			return cust
		}
	}
	return nil
}

func GetAllUsers() []*User {
	return AllUsers
}

func GetAccountsForUser(u *User) []*accService.Account {
	return u.Accounts
}

func NewAdmin(firstName, lastName, password string) (*User, error) {
	if firstName == "" || lastName == "" || password == "" {
		return nil, errEmptyString
	}

	tempCust := &User{
		UserID:    findUserID(),
		FirstName: firstName,
		LastName:  lastName,
		IsAdmin:   true,
		IsActive:  true,
		Password:  password,
		Accounts:  []*accService.Account{},
	}
	AllUsers = append(AllUsers, tempCust)
	return tempCust, nil
}

func NewUser(firstName, lastName, password string) (*User, error) {
	if firstName == "" || lastName == "" || password == "" {
		return nil, errEmptyString
	}

	tempCust := &User{
		UserID:    findUserID(),
		FirstName: firstName,
		LastName:  lastName,
		IsAdmin:   false,
		IsActive:  true,
		Password:  password,
		Accounts:  []*accService.Account{},
	}
	AllUsers = append(AllUsers, tempCust)
	return tempCust, nil
}

func GetUsers() ([]*User, error) {
	activeUsers := []*User{}
	for _, cust := range AllUsers {
		if cust.IsActive {
			activeUsers = append(activeUsers, cust)
		}
	}
	return activeUsers, nil
}

func UpdateUsers(userID int, parameter string, newValue interface{}) error {
	user := findUserByID(userID)
	if user == nil {
		return errNoUserFound
	}

	switch parameter {
	case "First Name":
		if value, ok := newValue.(string); ok {
			user.FirstName = value
		}
	case "Last Name":
		if value, ok := newValue.(string); ok {
			user.LastName = value
		}
	case "Admin Rights":
		if value, ok := newValue.(bool); ok {
			user.IsAdmin = value
		}
	case "Password":
		if value, ok := newValue.(string); ok {
			user.Password = value
		}
	default:
		return errors.New("invalid parameter entered")
	}
	return nil
}

func DeleteUsers(userID int) error {
	user := findUserByID(userID)
	if user == nil {
		return errNoUserFound
	}
	user.IsActive = false
	return nil
}

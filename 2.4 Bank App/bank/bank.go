// deprecating, using bankwithledger.go instead
package bank

import (
	"bankingApp/account"
	"errors"
)

// Bank represents a banking institution.
type bank struct {
	BankID       int
	FullName     string
	Abbreviation string
	IsActive     bool
	Accounts     []*account.Account
	// Ledger       map[string]float64
}

// allBanks stores all bank instances.
var allBanks = []*bank{}

// Error messages
var (
	errEmptyFullName     = errors.New("bank full name cannot be empty")
	errEmptyAbbreviation = errors.New("bank abbreviation cannot be empty")
	errInactiveBank      = errors.New("bank is inactive")
	errAccountNotFound   = errors.New("account not found")
	errInvalidParameter  = errors.New("invalid parameter entered")
)

// findBankID generates the next bank ID.
func findBankID() int {
	if len(allBanks) == 0 {
		return 0
	}
	return allBanks[len(allBanks)-1].BankID + 1
}

// NewBank creates a new bank and returns a pointer to it.
func newBank(fullName, abbreviation string) (*bank, error) {
	if err := validateBankName(fullName, abbreviation); err != nil {
		return nil, err
	}

	b := &bank{
		BankID:       findBankID(),
		FullName:     fullName,
		Abbreviation: abbreviation,
		IsActive:     true,
		Accounts:     []*account.Account{},
		// Ledger:       make(map[string]float64),
	}

	allBanks = append(allBanks, b)
	return b, nil
}

// validateBankName checks that bank names are not empty.
func validateBankName(fullName, abbreviation string) error {
	if fullName == "" {
		return errEmptyFullName
	}
	if abbreviation == "" {
		return errEmptyAbbreviation
	}
	return nil
}

// GetActivityStatus returns the bank's active status.
func (b *Bank) getActivityStatus() bool {
	return b.IsActive
}

// GetBankID returns the bank's ID.
func (b *Bank) getBankID() int {
	return b.BankID
}

func (b *Bank) getBankName() string {
	return b.FullName
}

func getallBanks() []*bank {
	return allBanks
}

// FindBankByID returns the bank's full name by its ID.
func findBankByID(bankID int) string {
	for _, bank := range allBanks {
		if bank.BankID == bankID {
			return bank.FullName
		}
	}
	return ""
}

func findBank(bankID int) *bank {
	for _, bank := range allBanks {
		if bank.BankID == bankID {
			return bank
		}
	}
	return nil
}

// AddAccount creates a new account for the bank with the specified initial payment.
func (b *bank) AddAccount(initialPayment float64) (*account.Account, error) {
	if !b.IsActive {
		return nil, errInactiveBank
	}

	accountNo := b.findAccountIDforBank()
	newAccount := account.NewAccount(accountNo, b.BankID, initialPayment)
	b.Accounts = append(b.Accounts, newAccount)
	return newAccount, nil
}

// findAccountIDforBank generates the next account ID for the bank.
func (b *bank) findAccountIDforBank() int {
	if len(b.Accounts) == 0 {
		return 0
	}
	return b.Accounts[len(b.Accounts)-1].AccountNo + 1
}

// GetAccountByID retrieves an account by its account number.
func (b *bank) GetAccountByID(accountNo int) (*account.Account, error) {
	for _, acc := range b.Accounts {
		if acc.AccountNo == accountNo {
			return acc, nil
		}
	}
	return nil, errAccountNotFound
}

// UpdateBankInformation updates the bank's information based on the provided parameter and value.
func (b *Bank) updateBankInformation(parameter, value string) error {
	switch parameter {
	case "Full Bank Name":
		if value == "" {
			return errEmptyFullName
		}
		b.FullName = value
	case "Abbreviation":
		if value == "" {
			return errEmptyAbbreviation
		}
		b.Abbreviation = value
	default:
		return errInvalidParameter
	}
	return nil
}

// RemoveBank marks the bank as inactive and removes all its accounts.
func (b *Bank) removeBank() {
	for _, acc := range b.Accounts {
		acc.RemoveAccount()
	}
	b.IsActive = false
}

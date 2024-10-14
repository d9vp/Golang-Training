package bank

import (
	"bankingApp/account"
)

type Bank struct {
	bank
	Ledger []*Ledger
}

// AllBanks stores all bank instances.
var AllBanks = []*Bank{}

// NewBank creates a new bank and returns a pointer to it.
func NewBank(fullName, abbreviation string) (*Bank, error) {
	if err := validateBankName(fullName, abbreviation); err != nil {
		return nil, err
	}

	b, _ := newBank(fullName, abbreviation)
	ban := &Bank{
		bank:   *b,
		Ledger: nil,
	}

	AllBanks = append(AllBanks, ban)
	return ban, nil
}

// validateBankName checks that bank names are not empty.
func ValidateBankName(fullName, abbreviation string) error {
	if fullName == "" {
		return errEmptyFullName
	}
	if abbreviation == "" {
		return errEmptyAbbreviation
	}
	return nil
}

// GetActivityStatus returns the bank's active status.
func (b *Bank) GetActivityStatus() bool {
	return b.bank.IsActive
}

// GetBankID returns the bank's ID.
func (b *Bank) GetBankID() int {
	return b.bank.BankID
}

func (b *Bank) GetBankName() string {
	return b.bank.FullName
}

func GetAllBanks() []*Bank {
	return AllBanks
}

// FindBankByID returns the bank's full name by its ID.
func FindBankByID(bankID int) string {
	for _, bank := range AllBanks {
		if bank.bank.BankID == bankID {
			return bank.FullName
		}
	}
	return ""
}

func FindBank(bankID int) *Bank {
	for _, bank := range AllBanks {
		if bank.bank.BankID == bankID {
			return bank
		}
	}
	return nil
}

// AddAccount creates a new account for the bank with the specified initial payment.
func (b *Bank) AddAccount(initialPayment float64) (*account.Account, error) {
	if !b.bank.IsActive {
		return nil, errInactiveBank
	}

	accountNo := b.bank.findAccountIDforBank()
	newAccount := account.NewAccount(accountNo, b.bank.BankID, initialPayment)
	b.bank.Accounts = append(b.bank.Accounts, newAccount)
	return newAccount, nil
}

// GetAccountByID retrieves an account by its account number.
func (b *Bank) GetAccountByID(accountNo int) (*account.Account, error) {
	for _, acc := range b.bank.Accounts {
		if acc.AccountNo == accountNo {
			return acc, nil
		}
	}
	return nil, errAccountNotFound
}

// UpdateBankInformation updates the bank's information based on the provided parameter and value.
func (b *Bank) UpdateBankInformation(parameter, value string) error {
	switch parameter {
	case "Full Bank Name":
		if value == "" {
			return errEmptyFullName
		}
		b.bank.FullName = value
	case "Abbreviation":
		if value == "" {
			return errEmptyAbbreviation
		}
		b.bank.Abbreviation = value
	default:
		return errInvalidParameter
	}
	return nil
}

// RemoveBank marks the bank as inactive and removes all its accounts.
func (b *Bank) RemoveBank() {
	for _, acc := range b.bank.Accounts {
		acc.RemoveAccount()
	}
	b.bank.IsActive = false
}

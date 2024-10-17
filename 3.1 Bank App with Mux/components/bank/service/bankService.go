package service

import (
	"errors"
	"user/components/accounts/service" // Import for account service interactions
)

var AllBanks = []*Bank{}

// Error messages
var (
	errBankNotFound     = errors.New("bank not found")
	errInvalidBankName  = errors.New("invalid bank name or abbreviation")
	errBankInactive     = errors.New("bank is inactive")
	errInvalidParameter = errors.New("invalid parameter for update")
)

// Bank struct
type Bank struct {
	ID           int                `json:"id"`
	FullName     string             `json:"fullName"`
	Abbreviation string             `json:"abbreviation"`
	IsActive     bool               `json:"isActive"`
	Accounts     []*service.Account `json:"accounts"`
	LedgerData   []*LedgerData      `json:"ledger"`
}

func (b *Bank) GetActivityStatus() bool {
	return b.IsActive
}

func (b *Bank) GetBankID() int {
	return b.ID
}

func (b *Bank) GetBankName() string {
	return b.FullName
}

func (b *Bank) GetBankAbbreviation() string {
	return b.Abbreviation
}

func GetAllBanks() []*Bank {
	return AllBanks
}

// Get a specific bank by ID
func GetBankByID(id int) (*Bank, error) {
	for _, bank := range AllBanks {
		if bank.ID == id {
			return bank, nil
		}
	}
	return nil, errBankNotFound
}

// Create a new bank
func CreateBank(bank *Bank) (*Bank, error) {
	if bank.FullName == "" || bank.Abbreviation == "" {
		return nil, errInvalidBankName
	}
	bank.ID = len(AllBanks) + 1
	bank.IsActive = true
	AllBanks = append(AllBanks, bank)
	return bank, nil
}

// Update bank information
func UpdateBank(id int, parameter string, newValue string) error {
	for i, bank := range AllBanks {
		if bank.ID == id {
			switch parameter {
			case "Full Name":
				if newValue == "" {
					return errors.New("full name cannot be empty")
				}
				AllBanks[i].FullName = newValue
			case "Abbreviation":
				if len(newValue) == 0 {
					return errors.New("abbreviation cannot be empty")
				} else if len(newValue) > 5 {
					return errors.New("abbreviation too long, must be under 5 characters")
				}
				AllBanks[i].Abbreviation = newValue

			default:
				return errInvalidParameter
			}
			return nil
		}
	}
	return errBankNotFound
}

// Remove a bank
func RemoveBank(bankID int) error {
	for i, bank := range AllBanks {
		if bank.ID == bankID {
			if !bank.IsActive {
				return errBankInactive
			}

			// Mark the bank as inactive and remove all associated accounts
			bank.IsActive = false
			for _, account := range bank.Accounts {
				service.DeleteAccount(account) // Call to account service to remove accounts
			}

			AllBanks[i] = bank
			return nil
		}
	}
	return errBankNotFound
}

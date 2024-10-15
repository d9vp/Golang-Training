package bank

import (
	"bankingApp/account"
	"errors"
)

type BankFunctions interface {
	AddAccount(initialPayment float64) (account.AccountFunctions, error)
	UpdateBankInformation(parameter, value string) error
	RemoveBank()
	findAccountIDforBank() int
	GetBankID() int
	GetBankName() string
	GetActivityStatus() bool
	GetBankAbbreviation() string
	AddToLedger(correspondingBankID int, amount float64)
	GetLedgerRecord()
}

type Bank struct {
	BankID       int
	FullName     string
	Abbreviation string
	IsActive     bool
	Accounts     []account.AccountFunctions
	Ledger       []*Ledger
}

var AllBanks = []BankFunctions{}

var (
	errEmptyFullName     = errors.New("bank full name cannot be empty")
	errEmptyAbbreviation = errors.New("bank abbreviation cannot be empty")
	errInactiveBank      = errors.New("bank is inactive")
	errAccountNotFound   = errors.New("account not found")
	errInvalidParameter  = errors.New("invalid parameter entered")
)

func findBankID() int {
	if len(AllBanks) == 0 {
		return 0
	}
	return AllBanks[len(AllBanks)-1].GetBankID() + 1
}

func NewBank(fullName, abbreviation string) (BankFunctions, error) {
	if err := validateBankName(fullName, abbreviation); err != nil {
		return nil, err
	}
	var ban BankFunctions = &Bank{
		BankID:       findBankID(),
		FullName:     fullName,
		Abbreviation: abbreviation,
		IsActive:     true,
		Accounts:     []account.AccountFunctions{},
		Ledger:       nil,
	}

	AllBanks = append(AllBanks, ban)
	return ban, nil
}

func validateBankName(fullName, abbreviation string) error {
	if fullName == "" {
		return errEmptyFullName
	}
	if abbreviation == "" {
		return errEmptyAbbreviation
	}
	return nil
}

func (b *Bank) GetActivityStatus() bool {
	return b.IsActive
}

func (b *Bank) GetBankID() int {
	return b.BankID
}

func (b *Bank) GetBankName() string {
	return b.FullName
}

func GetAllBanks() []BankFunctions {
	return AllBanks
}

func (b *Bank) GetBankAbbreviation() string {
	return b.Abbreviation
}

func FindBankByID(bankID int) string {
	for _, bank := range AllBanks {
		if bank.GetBankID() == bankID {
			return bank.GetBankName()
		}
	}
	return ""
}

func FindBank(bankID int) BankFunctions {
	for _, bank := range AllBanks {
		if bank.GetBankID() == bankID {
			return bank
		}
	}
	return nil
}
func (b *Bank) findAccountIDforBank() int {
	if len(b.Accounts) == 0 {
		return 0
	}
	return b.Accounts[len(b.Accounts)-1].GetAccountNumber() + 1
}

func (b *Bank) AddAccount(initialPayment float64) (account.AccountFunctions, error) {
	if !b.IsActive {
		return nil, errInactiveBank
	}

	accountNo := b.findAccountIDforBank()
	var newAccount account.AccountFunctions = account.NewAccount(accountNo, b.BankID, initialPayment)
	b.Accounts = append(b.Accounts, newAccount)
	return newAccount, nil
}

func (b *Bank) GetAccountByID(accountNo int) (account.AccountFunctions, error) {
	for _, acc := range b.Accounts {
		if acc.GetAccountNumber() == accountNo {
			return acc, nil
		}
	}
	return nil, errAccountNotFound
}

func (b *Bank) UpdateBankInformation(parameter, value string) error {
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

func (b *Bank) RemoveBank() {
	for _, acc := range b.Accounts {
		acc.RemoveAccount()
	}
	b.IsActive = false
}

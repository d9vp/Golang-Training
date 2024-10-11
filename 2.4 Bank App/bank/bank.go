package bank

import (
	"bankingApp/account"
	"errors"
	"fmt"
)

type Bank struct {
	BankID       int
	FullName     string
	Abbreviation string
	IsActive     bool
	Accounts     []*account.Account
}

var AllBanks = []*Bank{}

func findBankID() int {
	if len(AllBanks) == 0 {
		return 0
	} else {
		return AllBanks[len(AllBanks)-1].BankID + 1
	}
}

func (b *Bank) findAccountIDforBank() int {
	if b.IsActive {
		if len(b.Accounts) == 0 {
			return 0
		} else {
			return b.Accounts[len(b.Accounts)-1].AccountNo + 1
		}
	}
	return -1
}

func GetAllBanks() []*Bank {
	return AllBanks
}

func (bank *Bank) GetActivityStatus() bool {
	return bank.IsActive
}

func (bank *Bank) GetBankID() int {
	return bank.BankID
}

func GetBankNameFromID(bankID int) string {
	for _, bank := range AllBanks {
		if bank.BankID == bankID {
			return bank.FullName
		}
	}
	return ""
}

func validateBankName(fullName, abbreviation string) error {
	if fullName == "" {
		return errors.New("first name cannot be empty")
	}
	if abbreviation == "" {
		return errors.New("last name cannot be empty")
	}
	return nil

}

func NewBank(fullName, abbreviation string) *Bank {
	err := validateBankName(fullName, abbreviation)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	tempBank := &Bank{
		BankID:       findBankID(),
		FullName:     fullName,
		Abbreviation: abbreviation,
		IsActive:     true,
		Accounts:     []*account.Account{},
	}
	AllBanks = append(AllBanks, tempBank)
	return tempBank
}

func (b *Bank) AddAccount(firstPayment float64) *account.Account {
	if b.IsActive {
		accountNo := b.findAccountIDforBank()
		accountGenerated := account.NewAccount(accountNo, b.BankID, firstPayment)
		b.Accounts = append(b.Accounts, accountGenerated)
		return accountGenerated
	} else {
		fmt.Println("bank is inactive")
		return nil
	}
}

func (b *Bank) GetAccountByID(accountNo int) *account.Account {
	for _, c := range b.Accounts {
		if c.AccountNo == accountNo {
			return c
		}
	}
	fmt.Println("account not found")
	return nil
}

func (b *Bank) UpdateBankInformation(parameter, value string) {
	switch parameter {
	case "Full Bank Name":
		if value == "" {
			println("cannot enter empty string as a bank name")
			return
		}
		b.FullName = value
	case "Abbreviation":
		if value == "" {
			println("cannot enter empty string as a bank name abbreviation")
			return
		}
		b.Abbreviation = value
	default:
		println("Invalid parameter entered")
		return
	}
}

func (b *Bank) RemoveBank() {
	for _, a := range b.Accounts {
		a.RemoveAccount()
	}
	b.IsActive = false
}

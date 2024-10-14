package bank

import "fmt"

type Ledger struct {
	BankName string
	Amount   float64
}

func NewLedger(correspondingBankID int, amount float64) *Ledger {
	bankName := FindBankByID(correspondingBankID)
	return &Ledger{
		BankName: bankName,
		Amount:   amount,
	}
}

// AddToLedger records or updates a transaction in the bank's ledger.
func (b *Bank) AddToLedger(correspondingBankID int, amount float64) {
	for i, entry := range b.Ledger {
		if entry.BankName == FindBankByID(correspondingBankID) {
			b.Ledger[i].Amount += amount
			return
		}
	}

	tempLedger := NewLedger(correspondingBankID, amount)
	b.Ledger = append(b.Ledger, tempLedger)
}

func (b *Bank) GetLedgerRecord() {
	fmt.Printf("Ledger Record for %s:\n", b.FullName)
	fmt.Printf("%-20s | %-10s\n", "Corresponding Bank", "Amount")

	for _, entry := range b.Ledger {
		fmt.Printf("%-20s | %+.2f\n", entry.BankName, entry.Amount)
	}
}

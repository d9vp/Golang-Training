package bank

import "fmt"

type Ledger struct {
	BankName string
	Amount   float64
}

func NewLedger(correspondingBankID int, amount float64) *Ledger {
	bankName := findBankByID(correspondingBankID)
	return &Ledger{
		BankName: bankName,
		Amount:   amount,
	}
}

// AddToLedger records or updates a transaction in the bank's ledger.
func (b *Bank) AddToLedger(correspondingBankID int, amount float64) {
	// Check if the ledger already has an entry for the corresponding bank
	for i, entry := range b.Ledger {
		if entry.BankName == findBankByID(correspondingBankID) { // Check for existing bank name
			b.Ledger[i].Amount += amount // Increment the existing amount
			return
		}
	}

	// If no entry exists, create a new one
	tempLedger := NewLedger(correspondingBankID, amount)
	b.Ledger = append(b.Ledger, tempLedger)
}

// GetLedgerRecord displays the ledger records for the bank.
func (b *Bank) GetLedgerRecord() {
	fmt.Printf("Ledger Record for %s:\n", b.FullName)
	fmt.Printf("%-20s | %-10s\n", "Corresponding Bank", "Amount")

	// Display the ledger records
	for _, entry := range b.Ledger {
		fmt.Printf("%-20s | %+.2f\n", entry.BankName, entry.Amount)
	}
}
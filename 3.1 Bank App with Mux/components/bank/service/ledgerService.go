package service

type LedgerData struct {
	CorrespondingBankID int     `json:"corresponding_bank_id"`
	Amount              float64 `json:"amount"`
}

func AddToLedger(bankID int, corrBankID int, amount float64) error {
	b, _ := GetBankByID(bankID)
	if b == nil {
		return errBankNotFound
	}
	cb, _ := GetBankByID(corrBankID)
	if cb == nil {
		return errBankNotFound
	}
	flag := 0
	if b.LedgerData != nil {
		for _, ledger := range b.LedgerData {
			if ledger.CorrespondingBankID == corrBankID {
				flag = 1
				ledger.Amount += amount
			}
		}
	}
	if flag == 0 {
		ledger := &LedgerData{
			CorrespondingBankID: corrBankID,
			Amount:              amount,
		}
		b.LedgerData = append(b.LedgerData, ledger)
	}
	flag = 0
	if cb.LedgerData != nil {
		for _, ledger := range cb.LedgerData {
			if ledger.CorrespondingBankID == bankID {
				flag = 1
				ledger.Amount -= amount
			}
		}

	}
	if flag == 0 {
		ledger := &LedgerData{
			CorrespondingBankID: bankID,
			Amount:              -amount,
		}
		cb.LedgerData = append(cb.LedgerData, ledger)
	}

	return nil
}

func GetLedger(bankID int) ([]*LedgerData, error) {
	b, _ := GetBankByID(bankID)
	if b == nil {
		return nil, errBankNotFound
	}

	return b.LedgerData, nil
}

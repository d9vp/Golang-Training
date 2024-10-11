package account

import "fmt"

type Account struct {
	AccountNo int
	BankID    int
	Balance   float64
	IsActive  bool
}

func (a *Account) GetAccountNumber() int {
	return a.AccountNo
}
func (a *Account) GetBankID() int {
	return a.BankID
}
func (a *Account) GetActivityStatus() bool {
	return a.IsActive
}
func (a *Account) GetBalance() float64 {
	return a.Balance
}

func NewAccount(accountNo, bankID int, firstPayment float64) *Account {
	//validation carried out in the process of calling
	return &Account{
		AccountNo: accountNo,
		BankID:    bankID,
		Balance:   firstPayment,
		IsActive:  true,
	}
}

func (a *Account) Deposit(amount float64) {
	if amount <= 0 {
		fmt.Println("Deposit amount should be greater than 0")
		return
	}
	a.Balance += amount
	fmt.Printf("Deposited Rs. %.2f to account %d\n", amount, a.AccountNo)
}

func (a *Account) Withdraw(amount float64) {
	if amount > a.Balance {
		fmt.Println("Insufficient balance for withdrawal")
		return
	}
	a.Balance -= amount
	fmt.Printf("Withdrew Rs. %.2f from account %d\n", amount, a.AccountNo)
}

func (a *Account) Transfer(toAccount *Account, amount float64) {
	if amount > a.Balance {
		fmt.Println("Insufficient balance for transfer")
		return
	}
	a.Withdraw(amount)
	toAccount.Deposit(amount)
	fmt.Printf("Transferred Rs. %.2f from account %d to account %d\n", amount, a.AccountNo, toAccount.AccountNo)
}

func (a *Account) RemoveAccount() {
	a.IsActive = false
}

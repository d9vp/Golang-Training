package main

type Bank struct {
	BankID       int
	FullName     string
	Abbreviation string
	Accounts     []*Account
}

// CRUD

type Account struct {
	AccountNo int
	Bank      Bank
	Balance   float64
}

// other than CRUD
func (account *Account) Deposit(amount float64)
func (account *Account) Withdraw(amount float64)
func (account *Account) Transfer(toAccount *Account, amount float64)

type Customer struct {
	CustomerID   int
	FirstName    string
	LastName     string
	TotalBalance float64
	IsAdmin      bool
	Accounts     []*Account
}

// other than CRUD
func (customer *Customer) GetTotalBalance() float64
func (customer *Customer) DepositToAccount(bankID int, accountNo int, amount float64)
func (customer *Customer) WithdrawFromAccount(bankID int, accountNo int, amount float64)
func (customer *Customer) TransferBetweenAccounts(fromAccountNo int, fromBankID int, toAccountNo int, amount float64)

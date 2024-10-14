package user

import (
	"bankingApp/account"
	"bankingApp/bank"
	"errors"
	"fmt"
)

type Admin interface {
	NewUser(firstName, lastName string) User
	GetUsers()
	UpdateUsers(userID int, parameter string, newValue interface{})
	DeleteUsers(userID int)
	NewBank(fullName, abbreviation string) *User
	GetBanks()
	UpdateBank(bankID int, parameter string, newValue string)
	DeleteBank(bankID int)
	AddLedgerRecord(senderBankID, receiverBankID int, amount float64) *User
	GetLedgerRecord(bankID int) *User
}

type Customer interface {
	NewAccount(bankID int, initialPayment float64) *User
	GetAccounts()
	DeleteAccount(bankID, accountNo int)
	GetTotalBalance() float64
	DepositToAccount(accountNo int, bankID int, amount float64) *User
	WithdrawFromAccount(accountNo int, bankID int, amount float64) *User
	TransferFunds(fromAccountID, fromBankID, toAccountID, toBankID int, amount float64) *User
	TransferBetweenUsers(fromAccountNo, fromBankID, toCustID, toAccountNo, toBankID int, amount float64) *User
	GetPassbook(accountID int, bankID int) *User
}

type User struct {
	UserID    int
	FirstName string
	LastName  string
	IsAdmin   bool
	IsActive  bool
	Accounts  []*account.Account
}

var AllUsers = []*User{}

func (c *User) checkAdminAccess() error {
	if !c.IsAdmin {
		return errors.New("only admin access")
	}
	return nil
}

func (c *User) checkUserAccess() error {
	if c.IsAdmin {
		return errors.New("only user access")
	}
	return nil
}

func findUserID() int {
	if len(AllUsers) == 0 {
		return 0
	}
	return AllUsers[len(AllUsers)-1].UserID + 1
}

func (c *User) findAccount(accountNo, bankID int) *account.Account {
	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == accountNo && acc.GetBankID() == bankID && acc.IsActive {
			return acc
		}
	}
	return nil
}

func validateName(firstName, lastName string) error {
	if firstName == "" {
		return errors.New("first name cannot be empty")
	}
	if lastName == "" {
		return errors.New("last name cannot be empty")
	}
	return nil
}

func NewAdmin(firstName, lastName string) *User {
	if err := validateName(firstName, lastName); err != nil {
		fmt.Println(err)
		return nil
	}
	tempAdmin := &User{
		UserID:    findUserID(),
		FirstName: firstName,
		LastName:  lastName,
		IsAdmin:   true,
		IsActive:  true,
		Accounts:  nil,
	}
	AllUsers = append(AllUsers, tempAdmin)
	return tempAdmin
}

func (c *User) NewUser(firstName, lastName string) (*User, error) {
	if err := c.checkAdminAccess(); err != nil {
		return nil, err
	}
	if err := validateName(firstName, lastName); err != nil {
		return nil, err
	}
	tempCust := &User{
		UserID:    findUserID(),
		FirstName: firstName,
		LastName:  lastName,
		IsAdmin:   false,
		IsActive:  true,
		Accounts:  []*account.Account{},
	}
	AllUsers = append(AllUsers, tempCust)
	return tempCust, nil
}

func (c *User) NewBank(fullName, abbreviation string) *User {
	if err := c.checkAdminAccess(); err != nil {
		fmt.Println(err)
		return c
	}
	_, _ = bank.NewBank(fullName, abbreviation)
	return c
}

func (c *User) GetUsers() {
	if err := c.checkAdminAccess(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("All Users are listed as follows:")
	for _, cust := range AllUsers {
		if cust.IsActive {
			fmt.Printf("User ID:\t%d\nUser Name:\t%s %s\nAdmin Rights:\t%t\n",
				cust.UserID, cust.FirstName, cust.LastName, cust.IsAdmin)
		}
	}
}

func (c *User) findUserByID(userID int) *User {
	for _, cust := range AllUsers {
		if cust.UserID == userID && cust.IsActive {
			return cust
		}
	}
	return nil
}

func (c *User) UpdateUsers(userID int, parameter string, newValue interface{}) {
	if err := c.checkAdminAccess(); err != nil {
		fmt.Println(err)
		return
	}
	user := c.findUserByID(userID)
	if user == nil {
		fmt.Println("no such user id found")
		return
	}

	switch parameter {
	case "First Name":
		if value, ok := newValue.(string); ok {
			if err := validateName(value, user.LastName); err != nil {
				fmt.Println(err)
				return
			}
			user.FirstName = value
		}
	case "Last Name":
		if value, ok := newValue.(string); ok {
			if err := validateName(user.FirstName, value); err != nil {
				fmt.Println(err)
				return
			}
			user.LastName = value
		}
	case "Admin Rights":
		if value, ok := newValue.(bool); ok {
			user.IsAdmin = value
		}
	default:
		fmt.Println("Invalid parameter entered")
	}
}

func (c *User) DeleteUsers(userID int) {
	if !c.IsActive {
		fmt.Println("no such user found")
		return
	}
	if err := c.checkAdminAccess(); err != nil {
		fmt.Println(err)
		return
	}
	user := c.findUserByID(userID)
	if user == nil {
		fmt.Println("no such user id found")
		return
	}
	user.IsActive = false
	for _, acc := range user.Accounts {
		if acc.IsActive {
			acc.RemoveAccount()
		}
	}
}

func (c *User) GetBanks() {
	if err := c.checkAdminAccess(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("All Banks are listed as follows:")
	for _, bank := range bank.AllBanks {
		if bank.GetActivityStatus() {
			fmt.Printf("Bank ID:\t%d\nBank Name:\t%s\nAbbreviation:\t%s\n",
				bank.BankID, bank.FullName, bank.Abbreviation)
		}
	}
}

func (c *User) UpdateBank(bankID int, parameter string, newValue string) {
	if err := c.checkAdminAccess(); err != nil {
		fmt.Println(err)
		return
	}
	for _, bank := range bank.AllBanks {
		if bank.GetBankID() == bankID && bank.IsActive {
			bank.UpdateBankInformation(parameter, newValue)
			return
		}
	}
	fmt.Println("no such bank found")
}

func (c *User) DeleteBank(bankID int) {
	if !c.IsActive {
		fmt.Println("no such customer found")
		return
	}
	if err := c.checkAdminAccess(); err != nil {
		fmt.Println(err)
		return
	}
	for _, bank := range bank.AllBanks {
		if bank.GetBankID() == bankID && bank.GetActivityStatus() {
			bank.RemoveBank()
		}
	}
}

func (c *User) NewAccount(bankID int, initialPayment float64) *User {
	if err := c.checkUserAccess(); err != nil {
		fmt.Println(err)
		return c
	}
	if initialPayment < 1000.0 {
		fmt.Println("initial payment must be at least 1000")
		return c
	}
	// allBanks :=
	for _, ban := range bank.GetAllBanks() {
		if ban.BankID == bankID {
			acc, err := ban.AddAccount(initialPayment)
			if err != nil {
				fmt.Println(err)
				return c
			}
			c.Accounts = append(c.Accounts, acc)
			return c
		}
	}
	fmt.Println("incorrect bank id entered")
	return c
}

func (c *User) GetAccounts() {
	if !c.IsActive {
		fmt.Println("No such user found")
		return
	}
	if err := c.checkUserAccess(); err != nil {
		fmt.Println(err)
		return
	}
	if len(c.Accounts) == 0 {
		fmt.Println("No accounts found")
		return
	}

	for _, acc := range c.Accounts {
		if acc.GetActivityStatus() {
			fmt.Printf("Bank Name:\t%s\nAccount Number:\t%d\nBalance:\t%.2f\n",
				bank.FindBankByID(acc.GetBankID()), acc.GetAccountNumber(), acc.GetBalance())
		}
	}
	fmt.Printf("\nTotal Balance: %.2f\n", c.GetTotalBalance())
}

func (c *User) DeleteAccount(bankID, accountNo int) {
	if !c.IsActive {
		fmt.Println("No such user found")
		return
	}
	if err := c.checkUserAccess(); err != nil {
		fmt.Println(err)
		return
	}
	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == accountNo && acc.GetBankID() == bankID {
			acc.RemoveAccount()
			return
		}
	}
	fmt.Println("Account not found")
}

func (c *User) GetTotalBalance() float64 {
	if err := c.checkUserAccess(); err != nil {
		fmt.Println(err)
		return -1.0
	}

	total := 0.0
	for _, acc := range c.Accounts {
		if acc.GetActivityStatus() {
			total += acc.GetBalance()
		}
	}
	return total
}

func (c *User) DepositToAccount(accountNo int, bankID int, amount float64) *User {
	if err := c.checkUserAccess(); err != nil {
		fmt.Println(err)
		return c
	}
	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == accountNo && acc.GetBankID() == bankID && acc.GetActivityStatus() {
			acc.Deposit(amount)
			return c
		}
	}
	fmt.Println("account not found for this user")
	return c
}

func (c *User) WithdrawFromAccount(accountNo int, bankID int, amount float64) *User {
	if err := c.checkUserAccess(); err != nil {
		fmt.Println(err)
		return c
	}
	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == accountNo && acc.GetBankID() == bankID && acc.GetActivityStatus() {
			acc.Withdraw(amount)
			return c
		}
	}
	fmt.Println("account not found for this user")
	return c
}

func (c *User) TransferFunds(fromAccountID, fromBankID, toAccountID, toBankID int, amount float64) *User {
	if err := c.checkUserAccess(); err != nil {
		fmt.Println(err)
		return c
	}
	var fromAcc *account.Account
	var toAcc *account.Account

	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == fromAccountID && acc.GetBankID() == fromBankID {
			fromAcc = acc
		}
		if acc.GetAccountNumber() == toAccountID && acc.GetBankID() == toBankID {
			toAcc = acc
		}
	}
	if fromAcc != nil && toAcc != nil {
		fromAcc.Withdraw(amount)
		toAcc.Deposit(amount)
		return c
	}
	fmt.Println("account not found for this user")
	return c
}

func (c *User) TransferBetweenUsers(fromAccountNo, fromBankID, toCustID, toAccountNo, toBankID int, amount float64) *User {
	if err := c.checkUserAccess(); err != nil {
		fmt.Println(err)
		return c
	}

	fromAccount := c.findAccount(fromAccountNo, fromBankID)
	if fromAccount == nil {
		fmt.Println("Invalid sending account")
		return c
	}

	toUser := c.findUserByID(toCustID)
	if toUser == nil {
		fmt.Println("No such receiving user ID found")
		return c
	}

	toAccount := toUser.findAccount(toAccountNo, toBankID)
	if toAccount == nil {
		fmt.Println("Invalid receiving account")
		return c
	}

	fromAccount.Transfer(toAccount, amount)
	return c
}

func (c *User) GetPassbook(accountID int, bankID int) *User {
	if err := c.checkUserAccess(); err != nil {
		fmt.Println(err)
		return c
	}

	account := c.findAccount(accountID, bankID)
	if account == nil {
		fmt.Println("No such account found")
		return c
	}

	account.GetPassbook()
	return c
}

func (c *User) AddLedgerRecord(senderBankID, receiverBankID int, amount float64) *User {
	if err := c.checkAdminAccess(); err != nil {
		fmt.Println(err)
		return c
	}

	senderBank := bank.FindBank(senderBankID)
	if senderBank == nil {
		fmt.Println("No such sender bank ID found")
		return c
	}

	receiverBank := bank.FindBank(receiverBankID)
	if receiverBank == nil {
		fmt.Println("No such receiver bank ID found")
		return c
	}

	senderBank.AddToLedger(receiverBank.GetBankID(), amount)
	receiverBank.AddToLedger(senderBank.GetBankID(), -amount)
	return c
}

func (c *User) GetLedgerRecord(bankID int) *User {
	if err := c.checkAdminAccess(); err != nil {
		fmt.Println(err)
		return c
	}

	ban := bank.FindBank(bankID)
	if ban == nil {
		fmt.Println("No such bank found")
		return c
	}

	ban.GetLedgerRecord()
	return c
}

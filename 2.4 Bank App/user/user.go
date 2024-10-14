package user

import (
	"bankingApp/account"
	"bankingApp/bank"
	"errors"
	"fmt"
)

type Admin interface {
	NewUser(firstName, lastName string) (*User, error)
	GetUsers() error
	UpdateUsers(userID int, parameter string, newValue interface{}) error
	DeleteUsers(userID int) error
	NewBank(fullName, abbreviation string) (bank.BankFunctions, error)
	GetBanks() error
	UpdateBank(bankID int, parameter string, newValue string) error
	DeleteBank(bankID int) error
	AddLedgerRecord(senderBankID, receiverBankID int, amount float64) error
	GetLedgerRecord(bankID int) error
}

type Customer interface {
	NewAccount(bankID int, initialPayment float64) error
	GetAccounts() error
	DeleteAccount(bankID, accountNo int) error
	GetTotalBalance() float64
	DepositToAccount(accountNo int, bankID int, amount float64) error
	WithdrawFromAccount(accountNo int, bankID int, amount float64) error
	TransferFunds(fromAccountID, fromBankID, toAccountID, toBankID int, amount float64) error
	TransferBetweenUsers(fromAccountNo, fromBankID, toCustID, toAccountNo, toBankID int, amount float64) error
	GetPassbook(accountID int, bankID int) error
}

type User struct {
	UserID    int
	FirstName string
	LastName  string
	IsAdmin   bool
	IsActive  bool
	Accounts  []account.AccountFunctions
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

func (c *User) findAccount(accountNo, bankID int) account.AccountFunctions {
	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == accountNo && acc.GetBankID() == bankID && acc.GetActivityStatus() {
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
		Accounts:  []account.AccountFunctions{},
	}
	AllUsers = append(AllUsers, tempCust)
	return tempCust, nil
}

func (c *User) NewBank(fullName, abbreviation string) (bank.BankFunctions, error) {
	if err := c.checkAdminAccess(); err != nil {
		return nil, err
	}
	var bank1 bank.BankFunctions
	bank1, _ = bank.NewBank(fullName, abbreviation)
	return bank1, nil
}

func (c *User) GetUsers() error {
	if err := c.checkAdminAccess(); err != nil {
		return err
	}
	fmt.Println("All Users are listed as follows:")
	for _, cust := range AllUsers {
		if cust.IsActive {
			fmt.Printf("User ID:\t%d\nUser Name:\t%s %s\nAdmin Rights:\t%t\n",
				cust.UserID, cust.FirstName, cust.LastName, cust.IsAdmin)
		}
	}
	return nil
}

func (c *User) findUserByID(userID int) *User {
	for _, cust := range AllUsers {
		if cust.UserID == userID && cust.IsActive {
			return cust
		}
	}
	return nil
}

func (c *User) UpdateUsers(userID int, parameter string, newValue interface{}) error {
	if err := c.checkAdminAccess(); err != nil {
		return err
	}
	user := c.findUserByID(userID)
	if user == nil {
		return errors.New("no such user id found")
	}

	switch parameter {
	case "First Name":
		if value, ok := newValue.(string); ok {
			if err := validateName(value, user.LastName); err != nil {
				return err
			}
			user.FirstName = value
		}
	case "Last Name":
		if value, ok := newValue.(string); ok {
			if err := validateName(user.FirstName, value); err != nil {
				return err
			}
			user.LastName = value
		}
	case "Admin Rights":
		if value, ok := newValue.(bool); ok {
			user.IsAdmin = value
		}
	default:
		return errors.New("invalid parameter entered")
	}
	return nil
}

func (c *User) DeleteUsers(userID int) error {
	if !c.IsActive {
		return errors.New("no such user found")
	}
	if err := c.checkAdminAccess(); err != nil {
		fmt.Println(err)
		return err
	}
	user := c.findUserByID(userID)
	if user == nil {
		return errors.New("no such user id found")
	}
	user.IsActive = false
	for _, acc := range user.Accounts {
		if acc.GetActivityStatus() {
			acc.RemoveAccount()
		}
	}
	return nil
}

func (c *User) GetBanks() error {
	if err := c.checkAdminAccess(); err != nil {
		return err
	}
	fmt.Println("All Banks are listed as follows:")
	for _, bank := range bank.AllBanks {
		if bank.GetActivityStatus() {
			fmt.Printf("Bank ID:\t%d\nBank Name:\t%s\nAbbreviation:\t%s\n",
				bank.GetBankID(), bank.GetBankName(), bank.GetBankAbbreviation())
		}
	}
	return nil
}

func (c *User) UpdateBank(bankID int, parameter string, newValue string) error {
	if err := c.checkAdminAccess(); err != nil {
		return err
	}
	for _, bank := range bank.AllBanks {
		if bank.GetBankID() == bankID && bank.GetActivityStatus() {
			err := bank.UpdateBankInformation(parameter, newValue)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("no such bank found")
}

func (c *User) DeleteBank(bankID int) error {
	if !c.IsActive {
		return errors.New("no such customer found")
	}
	if err := c.checkAdminAccess(); err != nil {
		return err
	}
	for _, bank := range bank.AllBanks {
		if bank.GetBankID() == bankID && bank.GetActivityStatus() {
			bank.RemoveBank()
		}
	}
	return nil
}

func (c *User) NewAccount(bankID int, initialPayment float64) error {
	if err := c.checkUserAccess(); err != nil {
		return err
	}
	if initialPayment < 1000.0 {
		return errors.New("initial payment must be at least 1000")
	}
	// allBanks :=
	for _, ban := range bank.GetAllBanks() {
		if ban.GetBankID() == bankID {
			acc, err := ban.AddAccount(initialPayment)
			if err != nil {
				return err
			}
			c.Accounts = append(c.Accounts, acc)
			return nil
		}
	}
	return errors.New("incorrect bank id entered")
}

func (c *User) GetAccounts() error {
	if !c.IsActive {
		return errors.New("no such user found")
	}
	if err := c.checkUserAccess(); err != nil {
		fmt.Println(err)
		return err
	}
	if len(c.Accounts) == 0 {
		return errors.New("no accounts found")
	}

	for _, acc := range c.Accounts {
		if acc.GetActivityStatus() {
			fmt.Printf("Bank Name:\t%s\nAccount Number:\t%d\nBalance:\t%.2f\n",
				bank.FindBankByID(acc.GetBankID()), acc.GetAccountNumber(), acc.GetBalance())
		}
	}
	fmt.Printf("\nTotal Balance: %.2f\n", c.GetTotalBalance())
	return nil
}

func (c *User) DeleteAccount(bankID, accountNo int) error {
	if !c.IsActive {
		return errors.New("no such user found")
	}
	if err := c.checkUserAccess(); err != nil {
		return err
	}
	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == accountNo && acc.GetBankID() == bankID {
			acc.RemoveAccount()
			return nil
		}
	}
	return errors.New("account not found")
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

func (c *User) DepositToAccount(accountNo int, bankID int, amount float64) error {
	if err := c.checkUserAccess(); err != nil {
		return err
	}
	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == accountNo && acc.GetBankID() == bankID && acc.GetActivityStatus() {
			err := acc.Deposit(amount)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("account not found for this user")
}

func (c *User) WithdrawFromAccount(accountNo int, bankID int, amount float64) error {
	if err := c.checkUserAccess(); err != nil {
		return err
	}
	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == accountNo && acc.GetBankID() == bankID && acc.GetActivityStatus() {
			err := acc.Withdraw(amount)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("account not found for this user")
}

func (c *User) TransferFunds(fromAccountID, fromBankID, toAccountID, toBankID int, amount float64) error {
	if err := c.checkUserAccess(); err != nil {
		return err
	}
	var fromAcc account.AccountFunctions
	var toAcc account.AccountFunctions

	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == fromAccountID && acc.GetBankID() == fromBankID {
			fromAcc = acc
		}
		if acc.GetAccountNumber() == toAccountID && acc.GetBankID() == toBankID {
			toAcc = acc
		}
	}
	if fromAcc != nil && toAcc != nil {
		err1 := fromAcc.Withdraw(amount)
		if err1 != nil {
			return err1
		}
		err2 := toAcc.Deposit(amount)
		if err2 != nil {
			return err2
		}
		return nil
	}
	return errors.New("no such accounrt found")
}

func (c *User) TransferBetweenUsers(fromAccountNo, fromBankID, toCustID, toAccountNo, toBankID int, amount float64) error {
	if err := c.checkUserAccess(); err != nil {
		return err
	}

	fromAccount := c.findAccount(fromAccountNo, fromBankID)
	if fromAccount == nil {
		return errors.New("invalid sending account")
	}

	toUser := c.findUserByID(toCustID)
	if toUser == nil {
		return errors.New("no such receiving user id found")
	}

	toAccount := toUser.findAccount(toAccountNo, toBankID)
	if toAccount == nil {
		return errors.New("invalid receiving account")
	}

	fromAccount.Transfer(toAccount, amount)
	return nil
}

func (c *User) GetPassbook(accountID int, bankID int) error {
	if err := c.checkUserAccess(); err != nil {
		return err
	}

	account := c.findAccount(accountID, bankID)
	if account == nil {
		return errors.New("no such account found")
	}

	account.GetPassbook()
	return nil
}

func (c *User) AddLedgerRecord(senderBankID, receiverBankID int, amount float64) error {
	if err := c.checkAdminAccess(); err != nil {
		return err
	}

	senderBank := bank.FindBank(senderBankID)
	if senderBank == nil {
		return errors.New("no such sender bank id found")
	}

	receiverBank := bank.FindBank(receiverBankID)
	if receiverBank == nil {
		return errors.New("no such receiver bank id found")
	}

	senderBank.AddToLedger(receiverBank.GetBankID(), amount)
	receiverBank.AddToLedger(senderBank.GetBankID(), -amount)
	return nil
}

func (c *User) GetLedgerRecord(bankID int) error {
	if err := c.checkAdminAccess(); err != nil {
		return err
	}

	ban := bank.FindBank(bankID)
	if ban == nil {
		return errors.New("no such bank found")
	}

	ban.GetLedgerRecord()
	return nil
}

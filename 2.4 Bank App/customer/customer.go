package customer

import (
	"bankingApp/account"
	"bankingApp/bank"
	"errors"
	"fmt"
)

type Customer struct {
	CustomerID int
	FirstName  string
	LastName   string
	IsAdmin    bool
	IsActive   bool
	Accounts   []*account.Account
	// TotalBalance float64
}

var AllCustomers = []*Customer{}

func (c *Customer) onlyForAdmins() error {
	if c.IsAdmin {
		return nil
	} else {
		return errors.New("only admin access")
	}
}

func (c *Customer) onlyForCustomers() error {
	if !c.IsAdmin {
		return nil
	} else {
		return errors.New("only customer access")
	}
}

func findCustomerID() int {
	if len(AllCustomers) == 0 {
		return 0
	} else {
		return AllCustomers[len(AllCustomers)-1].CustomerID + 1
	}
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

func NewAdmin(firstName, lastName string) *Customer {
	err := validateName(firstName, lastName)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	tempAdmin := &Customer{
		CustomerID: findCustomerID(),
		FirstName:  firstName,
		LastName:   lastName,
		IsAdmin:    true,
		IsActive:   true,
		Accounts:   nil,
		// TotalBalance: 0,
	}
	AllCustomers = append(AllCustomers, tempAdmin)
	return tempAdmin
}

func (c *Customer) NewCustomer(firstName, lastName string) (*Customer, error) { //validate names
	err1 := c.onlyForAdmins()
	if err1 != nil {
		return nil, err1
	}
	err2 := validateName(firstName, lastName)
	if err2 != nil {
		return nil, err2
	}

	tempCust := &Customer{
		CustomerID: findCustomerID(),
		FirstName:  firstName,
		LastName:   lastName,
		IsAdmin:    false,
		IsActive:   true,
		Accounts:   []*account.Account{},
		// TotalBalance: 0,
	}
	AllCustomers = append(AllCustomers, tempCust)
	return tempCust, nil
}

func (c *Customer) NewBank(fullName, abbreviation string) *Customer {
	err := c.onlyForAdmins()
	if err != nil {
		fmt.Println(err)
		return c
	}
	_ = bank.NewBank(fullName, abbreviation)
	return c
}

func (c *Customer) GetCustomers() {
	err := c.onlyForAdmins()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("All Customers are listed as follows:")
	for _, cust := range AllCustomers {
		if cust.IsActive {
			fmt.Println("Customer ID:\t", cust.CustomerID)
			fmt.Println("Customer Name:\t", cust.FirstName+" "+cust.LastName)
			fmt.Println("Admin Rights:\t", cust.IsAdmin)
		}
	}
}

func (c *Customer) UpdateCustomers(customerID int, parameter string, newValue interface{}) {
	err := c.onlyForAdmins()
	if err != nil {
		fmt.Println(err)
		return
	}
	var customer *Customer = nil
	for _, cust := range AllCustomers {
		if cust.CustomerID == customerID && cust.IsActive {
			customer = cust
			break
		}
	}
	if customer == nil {
		fmt.Println("no such customer id found")
		return
	}

	switch parameter {
	case "First Name":
		value, ok := newValue.(string)
		if ok {
			err := validateName(value, customer.LastName)
			if err != nil {
				fmt.Println(err)
				return
			}
			customer.FirstName = value
		}

	case "Last Name":
		value, ok := newValue.(string)
		if ok {
			err := validateName(customer.FirstName, value)
			if err != nil {
				fmt.Println(err)
				return
			}
			customer.LastName = value
		}

	case "Admin Rights":
		value, ok := newValue.(bool)
		if ok {
			customer.IsAdmin = value
		}

	default:
		println("Invalid parameter entered")
		return
	}
}

func (c *Customer) DeleteCustomers(customerID int) {
	if !c.IsActive {
		fmt.Println("no such customer found")
		return
	}
	err := c.onlyForAdmins()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, cust := range AllCustomers {
		if cust.CustomerID == customerID {
			cust.IsActive = false
			for _, acc := range cust.Accounts {
				if acc.IsActive {
					acc.RemoveAccount()
				}
			}
		}
	}
}

func (c *Customer) GetBanks() {
	err := c.onlyForAdmins()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("All Banks are listed as follows:")
	for _, bank := range bank.AllBanks {
		if bank.GetActivityStatus() {
			fmt.Println("Bank ID:\t", bank.BankID)
			fmt.Println("Bank Name:\t", bank.FullName)
			fmt.Println("Abbreviation:\t", bank.Abbreviation)
		}
	}
}

func (c *Customer) UpdateBank(bankID int, parameter string, newValue string) {
	err := c.onlyForAdmins()
	if err != nil {
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

func (c *Customer) DeleteBank(bankID int) {
	if !c.IsActive {
		fmt.Println("no such customer found")
		return
	}
	err := c.onlyForAdmins()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, bank := range bank.AllBanks {
		if bank.GetBankID() == bankID && bank.GetActivityStatus() {
			bank.RemoveBank()
		}
	}
}

func (c *Customer) NewAccount(bankID int, initialPayment float64) *Customer {
	err := c.onlyForCustomers()
	if err != nil {
		fmt.Println(err)
		return c
	}
	if initialPayment < 1000.0 {
		fmt.Println("initial payment must be atleast 1000")
		return c
	}
	allBanks := bank.GetAllBanks()
	for _, ban := range allBanks {
		if ban.BankID == bankID {
			acc := ban.AddAccount(initialPayment)
			c.Accounts = append(c.Accounts, acc)
			return c
		}
	}
	fmt.Println("incorrect bank id entered")
	return c
}

func (c *Customer) GetAccounts() {
	if !c.IsActive {
		fmt.Println("No such user found")
		return
	}
	err := c.onlyForCustomers()
	if err != nil {
		fmt.Println(err)
		return
	}
	flag := 0

	for _, acc := range c.Accounts {
		if acc.GetActivityStatus() {
			flag = 1
			fmt.Println("Bank Name:\t", bank.GetBankNameFromID(acc.GetBankID()))
			fmt.Println("Account Number:\t", acc.GetAccountNumber())
			fmt.Println("Balance:\t", acc.GetBalance())
		}
	}
	if flag == 1 {
		fmt.Println("\nTotal Balance: ", c.GetTotalBalance())
	}
}

func (c *Customer) DeleteAccount(bankID, accountNo int) {
	if !c.IsActive {
		fmt.Println("No such user found")
		return
	}
	err := c.onlyForCustomers()
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == accountNo && acc.GetBankID() == bankID {
			acc.RemoveAccount()
		}
	}
}

func (c *Customer) GetTotalBalance() float64 {
	err := c.onlyForCustomers()
	if err != nil {
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

func (c *Customer) DepositToAccount(accountNo int, bankID int, amount float64) *Customer {
	err := c.onlyForCustomers()
	if err != nil {
		fmt.Println(err)
		return c
	}
	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == accountNo && acc.GetBankID() == bankID && acc.GetActivityStatus() {
			acc.Deposit(amount)
			return c
		}
	}
	fmt.Println(errors.New("account not found for this customer"))
	return c
}

func (c *Customer) WithdrawFromAccount(accountNo int, bankID int, amount float64) *Customer {
	err := c.onlyForCustomers()
	if err != nil {
		fmt.Println(err)
		return c
	}
	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == accountNo && acc.GetBankID() == bankID && acc.GetActivityStatus() {
			acc.Withdraw(amount)
			return c
		}
	}
	fmt.Println(errors.New("account not found for this customer"))
	return c
}

func (c *Customer) TransferBetweenOwnAccounts(fromAccountNo, fromBankID, toAccountNo, toBankID int, amount float64) *Customer {
	err := c.onlyForCustomers()
	if err != nil {
		fmt.Println(err)
		return c
	}
	var fromAccount, toAccount *account.Account
	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == fromAccountNo && acc.GetBankID() == fromBankID && acc.IsActive {
			fromAccount = acc
		}
		if acc.GetAccountNumber() == toAccountNo && acc.GetBankID() == toBankID && acc.IsActive {
			toAccount = acc
		}
	}
	if fromAccount != nil && toAccount != nil {
		fromAccount.Transfer(toAccount, amount)
	} else {
		fmt.Println(errors.New("invalid account numbers provided for transfer"))
	}
	return c
}

func (c *Customer) TransferBetweenCustomers(fromAccountNo, fromBankID, toCustID, toAccountNo, toBankID int, amount float64) *Customer {
	err := c.onlyForCustomers()
	if err != nil {
		fmt.Println(err)
		return c
	}
	var fromAccount, toAccount *account.Account
	for _, acc := range c.Accounts {
		if acc.GetAccountNumber() == fromAccountNo && acc.GetBankID() == fromBankID && acc.IsActive {
			fromAccount = acc
		}

	}
	var toCustomer *Customer = nil
	for _, cust := range AllCustomers {
		if cust.CustomerID == toCustID {
			toCustomer = cust
		}
	}
	if toCustomer == nil {
		fmt.Println("no such receiving customer id found")
		return c
	}

	for _, acc := range toCustomer.Accounts {
		if acc.GetAccountNumber() == toAccountNo && acc.GetBankID() == toBankID && acc.IsActive {
			toAccount = acc
		}
	}

	if fromAccount != nil && toAccount != nil {
		fromAccount.Transfer(toAccount, amount)
	} else {
		fmt.Println(errors.New("invalid account numbers provided for transfer"))
	}
	return c
}

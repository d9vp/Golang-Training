package main

import (
	"bankingApp/user"
	"fmt"
)

func main() {
	var admin1 user.Admin = user.NewAdmin("Admin", "Admin")
	var customer1, customer2, customer3 user.Customer
	customer1, _ = admin1.NewUser("Dev", "Patel")
	customer2, _ = admin1.NewUser("Krish", "Pandya")
	customer3, _ = admin1.NewUser("Vidhi", "Shah")

	_, _ = admin1.NewBank("Global Bank", "GB")
	_, _ = admin1.NewBank("National Bank", "NB")
	_, _ = admin1.NewBank("Indian Bank", "IB")

	_ = customer1.NewAccount(0, 10000)
	_ = customer1.NewAccount(1, 1500)
	_ = customer2.NewAccount(0, 5000)
	_ = customer3.NewAccount(1, 100000)

	admin1.GetUsers()
	// admin1.DeleteBank(0)
	customer1.GetAccounts()
	admin1.GetBanks()

	admin1.UpdateBank(0, "Full Bank Name", "HDFC")
	admin1.GetBanks()
	// customer1.DeleteAccount(0, 0)
	customer1.GetAccounts()
	customer1.TransferFunds(0, 1, 0, 0, 2000)
	customer1.GetAccounts()
	fmt.Println(customer1.GetTotalBalance())
	customer1.DepositToAccount(0, 0, 1000)
	customer1.WithdrawFromAccount(0, 0, 500.50)
	customer1.TransferBetweenUsers(0, 0, 2, 1, 0, 5500)
	customer1.GetAccounts()
	customer2.GetAccounts()
	// customer1.DeleteAccount(0, 0)
	customer1.GetPassbook(0, 0)

	admin1.AddLedgerRecord(1, 0, 2000)
	// admin1.GetLedgerRecord(1)
	admin1.AddLedgerRecord(2, 0, 300)
	// admin1.GetLedgerRecord(0)
	admin1.AddLedgerRecord(2, 1, 1300)
	// admin1.GetLedgerRecord(2)
	admin1.AddLedgerRecord(0, 1, 1000)
	admin1.GetLedgerRecord(0)
	admin1.GetLedgerRecord(1)
	admin1.GetLedgerRecord(2)

}

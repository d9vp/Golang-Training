package main

import (
	"bankingApp/customer"
)

func main() {
	admin1 := customer.NewAdmin("Admin", "Admin")
	customer1, _ := admin1.NewCustomer("Dev", "Patel")
	customer2, _ := admin1.NewCustomer("Krish", "Pandya")
	customer3, _ := admin1.NewCustomer("Vidhi", "Shah")
	_ = admin1.NewBank("Global Bank", "GB")
	_ = admin1.NewBank("National Bank", "NB")
	_ = customer1.NewAccount(0, 10000)
	_ = customer1.NewAccount(1, 1500)
	_ = customer2.NewAccount(0, 5000)
	_ = customer3.NewAccount(1, 100000)
	// admin1.GetCustomers()
	// admin1.DeleteBank(0)
	// customer1.GetAccounts()
	// admin1.GetBanks()

	// admin1.UpdateBank(0, "Full Bank Name", "HDFC")
	// admin1.GetBanks()
	// customer1.DeleteAccount(0, 0)
	// customer1.GetAccounts()
	// customer1.TransferBetweenOwnAccounts(0, 1, 0, 0, 2000)
	// customer1.GetAccounts()
	// fmt.Println(customer1.GetTotalBalance())

	customer1.TransferBetweenCustomers(0, 0, 2, 1, 0, 299)
	customer1.GetAccounts()
	customer2.GetAccounts()

}

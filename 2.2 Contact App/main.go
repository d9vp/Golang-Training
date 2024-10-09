package main

import (
	"contactApp/user"
)

func main() {
	admin1 := user.CreateAdminUser("Admin", "Patel")
	staff1 := admin1.CreateStaffUser("Dev", "Patel")
	staff2 := admin1.CreateStaffUser("Clone", "Patel")
	// fmt.Println(*admin1)
	// fmt.Println(*staff1)

	staff1, _ = staff1.CreateContact("Krish", "Pandya", 1)
	staff1, _ = staff1.CreateContact("Kush", "Desai", 1)
	staff1, _ = staff1.CreateContact("Vidhi", "Shah", 1)
	staff1, _ = staff1.CreateContact("Aditi", "Shah", 1)
	staff1, _ = staff1.CreateContact("Ketaki", "Bedekar", 1)
	// fmt.Println(*staff1)
	// user.GetContact(1, 0)

	_, _ = staff1.CreateContactInfo("email", "krishpandya@gmail.com", 1, 0)
	_, _ = staff1.CreateContactInfo("phone", "9282372838", 1, 0)
	_, _ = staff1.CreateContactInfo("email", "kushdesai@gmail.com", 1, 1)
	_, _ = staff1.CreateContactInfo("phone", "9726152415", 1, 2)
	_, _ = staff1.CreateContactInfo("email", "aditishah@gmail.com", 1, 3)
	_, _ = staff1.CreateContactInfo("email", "ketakibedekar@gmail.com", 1, 4)
	_, _ = staff1.CreateContactInfo("phone", "9172615442", 1, 4)
	// user.GetContactInfo(1, 0, 0)

	staff2, _ = staff2.CreateContact("Krish", "Pandya", 2)
	staff2, _ = staff2.CreateContact("Kush", "Desai", 2)
	staff2, _ = staff2.CreateContact("Vidhi", "Shah", 2)
	staff2, _ = staff2.CreateContact("Aditi", "Shah", 2)
	staff2, _ = staff2.CreateContact("Ketaki", "Bedekar", 2)
	// fmt.Println(*staff2)
	// user.GetContact(1, 0)

	_, _ = staff2.CreateContactInfo("email", "krishpandya@gmail.com", 2, 0)
	_, _ = staff2.CreateContactInfo("phone", "9282372838", 2, 0)
	_, _ = staff2.CreateContactInfo("email", "kushdesai@gmail.com", 2, 1)
	_, _ = staff2.CreateContactInfo("phone", "9726152415", 2, 2)
	_, _ = staff2.CreateContactInfo("email", "aditishah@gmail.com", 2, 3)
	_, _ = staff2.CreateContactInfo("email", "ketakibedekar@gmail.com", 2, 4)
	_, _ = staff2.CreateContactInfo("phone", "9172615442", 2, 4)

	// _, err := admin1.GetUser(0)
	// _, err = admin1.DeleteUser(0)
	// _, err := staff1.UpdateContact(1, 0, "Last Name", "Daksh")
	_, _ = staff1.UpdateContactInfo(1, 0, 0, "Contact Information Type", "phone")
	_, _ = staff1.UpdateContactInfo(1, 0, 0, "Contact Information Value", "9876543221")

	_, _ = staff1.GetContact(1, 0)
	// _, _ = staff1.GetContactInfo(1, 0, 0)

	// _, err = admin1.GetUser(1)
	// fmt.Println()

	// _, err = admin1.GetUser(2)
	// fmt.Println(err)
}

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

	staff1, _ = staff1.CreateContact("Krish", "Pandya")
	staff1, _ = staff1.CreateContact("Kush", "Desai")
	staff1, _ = staff1.CreateContact("Vidhi", "Shah")
	staff1, _ = staff1.CreateContact("Aditi", "Shah")
	staff1, _ = staff1.CreateContact("Ketaki", "Bedekar")
	// fmt.Println(*staff1)
	// user.GetContact(1, 0)

	_, _ = staff1.CreateContactInfo("email", "krishpandya@gmail.com", 0)
	_, _ = staff1.CreateContactInfo("phone", "9282372838", 0)
	_, _ = staff1.CreateContactInfo("email", "kushdesai@gmail.com", 1)
	_, _ = staff1.CreateContactInfo("phone", "9726152415", 2)
	_, _ = staff1.CreateContactInfo("email", "aditishah@gmail.com", 3)
	_, _ = staff1.CreateContactInfo("email", "ketakibedekar@gmail.com", 4)
	_, _ = staff1.CreateContactInfo("phone", "9172615442", 4)
	// user.GetContactInfo(1, 0, 0)

	staff2, _ = staff2.CreateContact("Krish", "Pandya")
	staff2, _ = staff2.CreateContact("Kush", "Desai")
	staff2, _ = staff2.CreateContact("Vidhi", "Shah")
	staff2, _ = staff2.CreateContact("Aditi", "Shah")
	staff2, _ = staff2.CreateContact("Ketaki", "Bedekar")
	// fmt.Println(*staff2)
	// user.GetContact(1, 0)

	_, _ = staff2.CreateContactInfo("email", "krishpandya@gmail.com", 0)
	_, _ = staff2.CreateContactInfo("phone", "9282372838", 0)
	_, _ = staff2.CreateContactInfo("email", "kushdesai@gmail.com", 1)
	_, _ = staff2.CreateContactInfo("phone", "9726152415", 2)
	_, _ = staff2.CreateContactInfo("email", "aditishah@gmail.com", 3)
	_, _ = staff2.CreateContactInfo("email", "ketakibedekar@gmail.com", 4)
	_, _ = staff2.CreateContactInfo("phone", "9172615442", 4)

	_, _ = admin1.GetUser(0)
	// admi1.GetUser
	// _, _ = admin1.DeleteUser(0)
	// _, err := staff1.UpdateContact(1, 0, "Last Name", "Daksh")
	// _, _ = staff1.UpdateContactInfo(1, 0, 0, "Contact Information Type", "phone")
	// _, _ = staff1.UpdateContactInfo(1, 0, 0, "Contact Information Value", "9876543221")

	// _, _ = admin1.UpdateUser("Last Name", "")
	// _, err := admin1.GetUser(0)
	// fmt.Println(err)

	// _, _ = staff1.GetContact(0)
	// _, _ = staff1.GetContactInfo(2, 0)

	// _, err := staff1.UpdateContactInfo(2, 0, "Contact Information Type", "phone")
	// fmt.Println(err)

	// _, err = staff1.UpdateContactInfo(2, 0, "Contact Information Value", "K&=@k.k")
	// fmt.Println(err)
	// _, _ = staff1.GetContactInfo(2, 0)

	// _, _ = admin1.DeleteUser()
	// fmt.Println(admin1.DeleteUser())
	// fmt.Println(admin1.GetUser(0))
	// admin1.GetUser(0)

	// fmt.Println(staff1.DeleteContactInfo(1, 0))
	// user, _ := staff1.GetContact(1)
	// fmt.Println(err)
	// add := (user.Contacts[1]).ContactInfo
	// fmt.Println(*add[0])
	// fmt.Println(user.Contacts[1])

	// _, _ = admin1.GetUser(1)
}

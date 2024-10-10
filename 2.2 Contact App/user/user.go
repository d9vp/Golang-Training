package user

import (
	"contactApp/contact"
	"errors"
	"fmt"
)

type User struct {
	UserID    int
	FirstName string
	LastName  string
	IsAdmin   bool
	IsActive  bool
	Contacts  []*contact.Contact
}

var AllUsers = []*User{}
var userIDs = 0

// /////////////////CREATION//////////////////////
func CreateAdminUser(firstName, lastName string) *User {
	err := validateUserInfo(firstName, lastName)
	if err != nil {
		panic(err)
	}
	admin1 := &User{
		UserID:    userIDs,
		FirstName: firstName,
		LastName:  lastName,
		IsAdmin:   true,
		IsActive:  true,
		Contacts:  nil,
	}
	userIDs++
	AllUsers = append(AllUsers, admin1)

	return admin1
}

func (user *User) CreateStaffUser(firstName, lastName string) *User {
	if !user.IsAdmin {
		panic("User needs to be an Admin to add new Staff.")
	}
	err := validateUserInfo(firstName, lastName)
	if err != nil {
		panic(err)
	}

	staff1 := &User{
		UserID:    userIDs,
		FirstName: firstName,
		LastName:  lastName,
		IsAdmin:   false,
		IsActive:  true,
		Contacts:  nil,
	}
	userIDs++
	AllUsers = append(AllUsers, staff1)

	return staff1
}

func (user *User) CreateContact(firstName, lastName string) (*User, error) {
	if user.IsAdmin {
		panic("Only staff can create contacts")
	}
	contactID := 0
	if len(user.Contacts) != 0 {
		contactID = user.Contacts[len(user.Contacts)-1].GetContactID()
		contactID++
	}
	tempContact := contact.CreateContact(firstName, lastName, contactID)
	user.Contacts = append(user.Contacts, tempContact)
	return user, nil
}

func (user *User) CreateContactInfo(contactInfoType, contactInfoValue string, contactID int) (*User, error) {
	if user.IsAdmin {
		panic("Only staff can add contact information")
	}

	for _, cont := range user.Contacts {
		if cont.GetContactID() == contactID {
			cont.CreateContactInfo(contactInfoType, contactInfoValue)
			return user, nil
		}
	}

	panic("no such user id found")

}

///////////////////READING//////////////////////

func (user *User) GetUser(userID int) (*User, error) {
	if !user.IsAdmin {
		return user, errors.New("only admins can view users")
	}
	for _, u := range AllUsers {
		// fmt.Println(u.UserID, userID)
		if u.UserID == userID && u.IsActive {
			fmt.Println("User ID: ", u.UserID)
			fmt.Println("First Name: ", u.FirstName)
			fmt.Println("Last Name: ", u.LastName)
			fmt.Println("Is Admin?: ", u.IsAdmin)
			fmt.Println("Is Active?: ", u.IsActive)
			return user, nil
		}
	}
	return user, errors.New("no such user id found")
}

func (user *User) GetContact(contactID int) (*User, error) {
	if user.IsAdmin {
		return user, errors.New("only staff can view contacts")
	}
	flag := 0
	if user.IsActive {
		for _, cont := range user.Contacts {
			if cont.GetContactID() == contactID && cont.GetActivityStatus() {
				flag = 1
				fmt.Println(*cont)
				for _, contInfo := range cont.ContactInfo {
					err := cont.GetContactInfo(contInfo.ContactInfoID)
					if err != nil {
						return user, err
					}
				}
			}
		}
		if flag == 0 {
			return user, errors.New("no such contact id found")
		} else {
			return user, nil
		}
	} else {
		return user, errors.New("no such user found")
	}
}

func (user *User) GetContactInfo(contactID, contactInfoID int) (*User, error) {
	if user.IsAdmin {
		return user, errors.New("only staff can view contact information")
	}
	flag := 0
	for _, cont := range user.Contacts {
		if cont.GetContactID() == contactID && cont.GetActivityStatus() {
			flag = 1
			err := cont.GetContactInfo(contactInfoID)
			if err != nil {
				return user, err
			}
		}
	}
	if flag == 0 {
		return user, errors.New("no such contact id found")
	} else {
		return user, nil
	}
}

// func (user *User) GetEverythingAllAtOnce(userID int) (*User, error) {
// 	flag := 0
// 	if !user.IsAdmin {
// 		return user, errors.New("only admins can view users")
// 	}
// 	for _, userN := range AllUsers {
// 		if userN.UserID == userID && userN.IsActive {
// 			flag = 1
// 			fmt.Println(*userN)
// 			for _, cont := range userN.Contacts {
// 				err := contact.GetContact(cont.ContactID, userN.Contacts)
// 				if err != nil {
// 					return user, err
// 				}
// 			}
// 		}
// 	}
// 	if flag == 0 {
// 		return user, errors.New("no such user id found")
// 	} else {
// 		return user, nil
// 	}
// }

//////////////UPDATION////////////

func (user *User) UpdateUser(parameter string, newValue interface{}) (*User, error) {
	if !user.IsAdmin {
		return user, errors.New("only admins can update user information")
	}
	if !user.IsActive {
		return user, errors.New("no such user found")
	}
	switch parameter {
	case "First Name":
		_, err := user.updateFirstName(newValue)
		if err != nil {
			return user, err
		} else {
			fmt.Println("Update successful!")
		}

	case "Last Name":
		_, err := user.updateLastName(newValue)
		if err != nil {
			return user, err
		} else {
			fmt.Println("Update successful!")
		}

	default:
		return user, errors.New("no such parameter found")

	}
	return user, nil
}

func (user *User) updateFirstName(newValue interface{}) (*User, error) {
	if value, ok := newValue.(string); ok {
		if newValue == "" {
			return user, errors.New("first name cannot be empty")
		}
		user.FirstName = value
		return user, nil
	} else {
		return user, errors.New("invalid type for first name expected a string")
	}

}

func (user *User) updateLastName(newValue interface{}) (*User, error) {
	if value, ok := newValue.(string); ok {
		if newValue == "" {
			return user, errors.New("last name cannot be empty")
		}
		user.LastName = value
		return user, nil
	} else {
		return user, errors.New("invalid type for last name expected a string")
	}

}

func (user *User) UpdateContact(contactID int, parameter string, newValue interface{}) (*User, error) {
	if user.IsAdmin {
		return user, errors.New("only staff can update contacts")
	}
	if !user.IsActive {
		return user, errors.New("no such user found")
	}
	for _, cont := range user.Contacts {
		if cont.GetContactID() == contactID && cont.GetActivityStatus() {
			err := cont.UpdateContact(parameter, newValue)
			if err == nil {
				return user, nil
			} else {
				return user, err
			}
		}
	}
	return user, errors.New("no such user id found")
}

func (user *User) UpdateContactInfo(contactID, contactInfoID int, parameter string, newValue interface{}) (*User, error) {
	if user.IsAdmin {
		return user, errors.New("only staff can update contacts")
	}
	if !user.IsActive {
		return user, errors.New("no such user found")
	}
	for _, cont := range user.Contacts {
		if cont.GetContactID() == contactID {
			err := cont.UpdateContactInfo(contactInfoID, parameter, newValue)
			if err == nil {
				return user, nil
			} else {
				return user, err
			}
		}
	}
	return user, errors.New("no such contact id found")
}

///////////DELETION/////////////

func (user *User) DeleteUser() (*User, error) {
	flag := 0
	if !user.IsAdmin {
		return user, errors.New("only admins can delete users")
	}

	if user.IsActive {
		user.IsActive = false
		flag = 1
		fmt.Println("Delete successful!")
	}

	if flag == 1 {
		return user, nil
	} else {
		return user, errors.New("no such user id found")
	}
}

func (user *User) DeleteContact(contactID int) (*User, error) {
	if user.IsAdmin {
		return user, errors.New("only staff can delete contacts")
	}
	if !user.IsActive {
		return user, errors.New("no such user found")
	}
	for _, contact := range user.Contacts {
		if contact.GetContactID() == contactID && contact.GetActivityStatus() {
			contact.IsActive = false
			return user, nil
		}
	}
	return user, errors.New("no such user id found")
}

func (user *User) DeleteContactInfo(contactID, contactInfoID int) (*User, error) {
	if user.IsAdmin {
		return user, errors.New("only staff can delete contacts")
	}
	if !user.IsActive {
		return user, errors.New("no such user found")
	}
	for _, contact := range user.Contacts {
		if contact.GetContactID() == contactID && contact.GetActivityStatus() {
			err := contact.DeleteContactInfo(contactInfoID)
			if err == nil {
				return user, nil
			} else {
				return user, err
			}
		}
	}
	return user, errors.New("no such contact if found")
}

func validateUserInfo(firstName, lastName string) error {
	if firstName == "" || lastName == "" {
		return errors.New("first name or last name cannot be empty")
	}
	return nil
}

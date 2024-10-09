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

func (user *User) CreateContact(firstName, lastName string, userID int) (*User, error) {
	if user.IsAdmin {
		panic("Only staff can create contacts")
	}
	contactID := 0
	for _, user := range AllUsers {
		if user.UserID == userID {
			if len(user.Contacts) != 0 {
				contactID = user.Contacts[len(user.Contacts)-1].ContactID
				contactID++
			}
			tempContact := contact.CreateContact(firstName, lastName, contactID)
			user.Contacts = append(user.Contacts, tempContact)
			return user, nil
		}
	}
	return user, errors.New("no user with given id found")
}

func (user *User) CreateContactInfo(contactInfoType, contactInfoValue string, userID, contactID int) (*User, error) {
	if user.IsAdmin {
		panic("Only staff can add contact information")
	}
	for _, user := range AllUsers {
		if user.UserID == userID {
			contact.CreateContactInfo(contactInfoType, contactInfoValue, contactID, user.Contacts)
		}
	}
	return user, errors.New("no user with given id found")

}

///////////////////READING//////////////////////

func (user *User) GetUser(userID int) (*User, error) {
	if !user.IsAdmin {
		return user, errors.New("only admins can view users")
	}
	for _, userN := range AllUsers {
		if userN.UserID == userID && userN.IsActive {
			// fmt.Println(*userN)
			fmt.Println("User ID: ", userN.UserID)
			fmt.Println("First Name: ", userN.FirstName)
			fmt.Println("Last Name: ", userN.LastName)
			fmt.Println("Is Admin?: ", userN.IsAdmin)
			fmt.Println("Is Active?: ", userN.IsActive)
			return user, nil
		}
	}
	return user, errors.New("no such user id found")
}

func (user *User) GetContact(userID, contactID int) (*User, error) {
	flag := 0
	if user.IsAdmin {
		return user, errors.New("only staff can view contacts")
	}
	for _, userN := range AllUsers {
		if userN.UserID == userID && userN.IsActive {
			flag = 1
			err := contact.GetContact(contactID, userN.Contacts)
			if err != nil {
				return user, err
			}
		}
	}
	if flag == 0 {
		return user, errors.New("no such user id found")
	} else {
		return user, nil
	}
}

func (user *User) GetContactInfo(userID, contactID, contactInfoID int) (*User, error) {
	if user.IsAdmin {
		return user, errors.New("only staff can view contact information")
	}
	for _, user := range AllUsers {
		if user.UserID == userID && user.IsActive {
			err := contact.GetContactInfo(contactID, contactInfoID, user.Contacts)
			return user, err
		}
	}
	return user, errors.New("no such user id found")
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

func (user *User) UpdateUser(userID int, parameter string, newValue interface{}) (*User, error) {
	if !user.IsAdmin {
		return user, errors.New("only admins can update user information")
	}
	for _, userN := range AllUsers {
		if userN.UserID == userID && userN.IsActive {
			switch parameter {
			case "First Name":
				_, err := userN.updateFirstName(newValue)
				if err != nil {
					panic(err)
				} else {
					fmt.Println("Update successful!")
				}

			case "Last Name":
				_, err := userN.updateLastName(newValue)
				if err != nil {
					panic(err)
				} else {
					fmt.Println("Update successful!")
				}

			default:
				return user, errors.New("no such parameter found")

			}
		}
	}
	return user, errors.New("no such user id found")
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

func (user *User) UpdateContact(userID, contactID int, parameter string, newValue interface{}) (*User, error) {
	if user.IsAdmin {
		return user, errors.New("only staff can update contacts")
	}
	for _, userN := range AllUsers {
		if userN.UserID == userID {
			err := contact.UpdateContact(contactID, parameter, newValue, userN.Contacts)
			if err == nil {
				return user, nil
			} else {
				return user, err
			}
		}
	}
	return user, errors.New("no such user id found")
}

func (user *User) UpdateContactInfo(userID, contactID, contactInfoID int, parameter string, newValue interface{}) (*User, error) {
	if user.IsAdmin {
		return user, errors.New("only staff can update contacts")
	}
	for _, userN := range AllUsers {
		if userN.UserID == userID {
			err := contact.UpdateContactInfo(contactID, contactInfoID, parameter, newValue, userN.Contacts)
			if err == nil {
				return user, nil
			} else {
				return user, err
			}
		}
	}
	return user, errors.New("no such user id found")
}

///////////DELETION/////////////

func (user *User) DeleteUser(userID int) (*User, error) {
	flag := 0
	if !user.IsAdmin {
		return user, errors.New("only admins can delete users")
	}
	for _, userN := range AllUsers {
		if userN.UserID == userID && userN.IsActive {
			userN.IsActive = false
			flag = 1
		}
	}
	if flag == 1 {
		return user, nil
	} else {
		return user, errors.New("no such user id found")
	}
}

func (user *User) DeleteContact(userID, contactID int) (*User, error) {
	if user.IsAdmin {
		return user, errors.New("only staff can delete contacts")
	}
	for _, userN := range AllUsers {
		if userN.UserID == userID && userN.IsActive {
			err := contact.DeleteContact(contactID, userN.Contacts)
			if err == nil {
				return user, nil
			} else {
				return user, err
			}
		}
	}
	return user, errors.New("no such user id found")
}

func (user *User) DeleteContactInfo(userID, contactID, contactInfoID int) (*User, error) {
	if user.IsAdmin {
		return user, errors.New("only staff can delete contacts")
	}
	for _, userN := range AllUsers {
		if userN.UserID == userID && userN.IsActive {
			err := contact.DeleteContactInfo(contactID, contactInfoID, userN.Contacts)
			if err == nil {
				return user, nil
			} else {
				return user, err
			}
		}
	}
	return user, errors.New("no such user id found")
}

func validateUserInfo(firstName, lastName string) error {
	if firstName == "" || lastName == "" {
		return errors.New("first name or last name cannot be empty")
	}
	return nil
}

package contact

import (
	"contactApp/contactInfo"
	"errors"
	"fmt"
)

type Contact struct {
	ContactID   int
	FirstName   string
	LastName    string
	IsActive    bool
	ContactInfo []*contactInfo.ContactInformation
}

func (contact *Contact) GetContactID() int {
	return contact.ContactID
}

func (contact *Contact) GetActivityStatus() bool {
	return contact.IsActive
}

// CREATION
func CreateContact(firstName, lastName string, contactID int) *Contact {
	err := validateUserInfo(firstName, lastName)
	if err != nil {
		panic(err)
	}

	tempContact := &Contact{
		ContactID:   contactID,
		FirstName:   firstName,
		LastName:    lastName,
		IsActive:    true,
		ContactInfo: nil,
	}
	return tempContact
}

func (contact *Contact) CreateContactInfo(contactInfoType, contactInfoValue string) {
	contactInfoID := 0

	if len(contact.ContactInfo) != 0 {
		contactInfoID = contact.ContactInfo[len(contact.ContactInfo)-1].ContactInfoID
		contactInfoID++
	}
	tempContactInfo := contactInfo.CreateContactInfoForContactID(contactInfoType, contactInfoValue, contactInfoID)
	contact.ContactInfo = append(contact.ContactInfo, tempContactInfo)

}

//READING

func (contact *Contact) GetContactInfo(contactInfoID int) error {
	if !contact.IsActive {
		return errors.New("no such contact found")
	}
	flag := 0
	for _, contactInfo := range contact.ContactInfo {
		if contactInfo.GetContactInfoID() == contactInfoID && contactInfo.GetContactInfoStatus() {
			flag = 1
			fmt.Println(*contactInfo)
		}
	}
	if flag == 0 {
		return errors.New("no such contact id found")
	} else {
		return nil
	}
}

//UPDATION

func (contact *Contact) UpdateContact(parameter string, newValue interface{}) error {
	switch parameter {
	case "First Name":
		_, err := contact.updateFirstName(newValue)
		if err != nil {
			return err
		} else {
			fmt.Println("Update successful!")
			return nil
		}

	case "Last Name":
		_, err := contact.updateLastName(newValue)
		if err != nil {
			return err
		} else {
			fmt.Println("Update successful!")
			return nil
		}

	default:
		return errors.New("no such parameter found")
	}
}

func (contact *Contact) updateFirstName(newValue interface{}) (*Contact, error) {
	if value, ok := newValue.(string); ok {
		if newValue == "" {
			return contact, errors.New("first name cannot be empty")
		}
		contact.FirstName = value
		return contact, nil
	} else {
		return contact, errors.New("invalid type for first name expected a string")
	}
}

func (contact *Contact) updateLastName(newValue interface{}) (*Contact, error) {
	if value, ok := newValue.(string); ok {
		if newValue == "" {
			return contact, errors.New("last name cannot be empty")
		}
		contact.LastName = value
		return contact, nil
	} else {
		return contact, errors.New("invalid type for last name expected a string")
	}
}

func (contact *Contact) UpdateContactInfo(contactInfoID int, parameter string, newValue interface{}) error {
	flag := 0

	for _, contactInfo := range contact.ContactInfo {
		if contactInfo.GetContactInfoID() == contactInfoID && contactInfo.GetContactInfoStatus() {
			flag = 1
			err := contactInfo.UpdateContactInfo(parameter, newValue)
			if err != nil {
				return err
			}
		}
	}
	if flag == 0 {
		return errors.New("no such contact information id found")
	} else {
		return nil
	}
}

//DELETION

func (contact *Contact) DeleteContactInfo(contactInfoID int) error {
	if !contact.IsActive {
		return errors.New("no such contact found")
	}
	for _, contactInfo := range contact.ContactInfo {
		if contactInfo.GetContactInfoID() == contactInfoID && contactInfo.GetContactInfoStatus() {
			contactInfo.IsActive = false
			return nil
		}
	}
	return errors.New("no such contact information id found")
}

func validateUserInfo(firstName, lastName string) error {
	if firstName == "" || lastName == "" {
		return errors.New("first name or last name cannot be empty")
	}
	return nil
}

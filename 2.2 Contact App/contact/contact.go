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

func CreateContactInfo(contactInfoType, contactInfoValue string, contactID int, contacts []*Contact) {
	contactInfoID := 0

	for _, contact := range contacts {
		if contact.ContactID == contactID {
			if len(contact.ContactInfo) != 0 {
				contactInfoID = contact.ContactInfo[len(contact.ContactInfo)-1].ContactInfoID
				contactInfoID++
			}
			tempContactInfo := contactInfo.CreateContactInfoForContactID(contactInfoType, contactInfoValue, contactInfoID)
			contact.ContactInfo = append(contact.ContactInfo, tempContactInfo)

			return
		}
	}
	panic("No such contact ID found.")
}

//READING

func GetContact(contactID int, contacts []*Contact) error {
	flag := 0
	for _, contact := range contacts {
		if contact.ContactID == contactID && contact.IsActive {
			flag = 1
			fmt.Println(*contact)
			for _, contInfo := range contact.ContactInfo {
				err := contactInfo.GetContactInfo(contInfo.ContactInfoID, contact.ContactInfo)
				if err != nil {
					return err
				}
			}
		}
	}
	if flag == 0 {
		return errors.New("no such contact id found")
	} else {
		return nil
	}
}

func GetContactInfo(contactID, contactInfoID int, contacts []*Contact) error {
	flag := 0
	for _, contact := range contacts {
		if contact.ContactID == contactID && contact.IsActive {
			flag = 1
			err := contactInfo.GetContactInfo(contactInfoID, contact.ContactInfo)
			if err != nil {
				return err
			}
		}
	}
	if flag == 0 {
		return errors.New("no such contact id found")
	} else {
		return nil
	}
}

//UPDATION

func UpdateContact(contactID int, parameter string, newValue interface{}, contacts []*Contact) error {
	for _, contact := range contacts {
		if contact.ContactID == contactID && contact.IsActive {
			switch parameter {
			case "First Name":
				_, err := contact.updateFirstName(newValue)
				if err != nil {
					panic(err)
				} else {
					fmt.Println("Update successful!")
					return nil
				}

			case "Last Name":
				_, err := contact.updateLastName(newValue)
				if err != nil {
					panic(err)
				} else {
					fmt.Println("Update successful!")
					return nil
				}

			default:
				return errors.New("no such parameter found")
			}
		}
	}
	return errors.New("no such contact id found")
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

func UpdateContactInfo(contactID, contactInfoID int, parameter string, newValue interface{}, contacts []*Contact) error {
	flag := 0
	for _, contact := range contacts {
		if contact.ContactID == contactID && contact.IsActive {
			flag = 1
			err := contactInfo.UpdateContactInfo(contactInfoID, parameter, newValue, contact.ContactInfo)
			if err != nil {
				return err
			}
		}
	}
	if flag == 0 {
		return errors.New("no such contact id found")
	} else {
		return nil
	}
}

//DELETION

func DeleteContact(contactID int, contacts []*Contact) error {
	flag := 0
	for _, contact := range contacts {
		if contact.ContactID == contactID && contact.IsActive {
			contact.IsActive = false
			flag = 1
		}
	}
	if flag == 1 {
		return nil
	} else {
		return errors.New("no such contact id found")
	}
}

func DeleteContactInfo(contactID, contactInfoID int, contacts []*Contact) error {
	for _, contact := range contacts {
		if contact.ContactID == contactID && contact.IsActive {
			err := contactInfo.DeleteContactInfo(contactInfoID, contact.ContactInfo)
			if err == nil {
				return nil
			} else {
				return err
			}
		}
	}
	return errors.New("no such contact id found")
}

func validateUserInfo(firstName, lastName string) error {
	if firstName == "" || lastName == "" {
		return errors.New("first name or last name cannot be empty")
	}
	return nil
}

package contactInfo

import (
	"errors"
	"fmt"
)

type ContactInformation struct {
	ContactInfoID    int
	ContactInfoType  string
	ContactInfoValue string
	IsActive         bool
}

func (contactinfo *ContactInformation) GetContactInfoID() int {
	return contactinfo.ContactInfoID
}

func (contactinfo *ContactInformation) GetContactInfoStatus() bool {
	return contactinfo.IsActive
}

//CREATING

func CreateContactInfoForContactID(contactInfoType, contactInfoValue string, contactInfoID int) *ContactInformation {
	err := validateContactInfo(contactInfoType, contactInfoValue)
	if err != nil {
		panic(err)
	}
	tempContactInfo := &ContactInformation{
		ContactInfoID:    contactInfoID,
		ContactInfoType:  contactInfoType,
		ContactInfoValue: contactInfoValue,
		IsActive:         true,
	}
	return tempContactInfo
}

// READING
func GetContactInfo(contactInfoID int, contactInfos []*ContactInformation) error {
	flag := 0
	for _, contactInfo := range contactInfos {
		if contactInfo.ContactInfoID == contactInfoID && contactInfo.IsActive {
			flag = 1
			fmt.Println(*contactInfo)
		}
	}
	if flag == 0 {
		return errors.New("no such contact information id found")
	} else {
		return nil
	}
}

//UPDATING

func (contactInfo *ContactInformation) UpdateContactInfo(parameter string, newValue interface{}) error {

	switch parameter {
	case "Contact Information Type":
		_, err := contactInfo.updateContactInfoType(newValue)
		if err != nil {
			return (err)
		} else {
			fmt.Println("Update successful!")
			return nil
		}

	case "Contact Information Value":
		_, err := contactInfo.updateContactInfoValue(newValue)
		if err != nil {
			return (err)
		} else {
			fmt.Println("Update successful!")
			return nil
		}

	default:
		return errors.New("no such parameter found")
	}
}

func (contactInfo *ContactInformation) updateContactInfoType(newValue interface{}) (*ContactInformation, error) {
	if value, ok := newValue.(string); ok {
		if newValue == "" {
			return contactInfo, errors.New("contact information type cannot be empty")
		}
		if value == "phone" && len(contactInfo.ContactInfoValue) != 10 {
			return contactInfo, errors.New("phone number must be 10 digits")
		}
		contactInfo.ContactInfoType = value
		return contactInfo, nil
	} else {
		return contactInfo, errors.New("invalid contact information type expected a string")
	}
}

func (contactInfo *ContactInformation) updateContactInfoValue(newValue interface{}) (*ContactInformation, error) {
	if value, ok := newValue.(string); ok {
		if newValue == "" {
			return contactInfo, errors.New("contact information value cannot be empty")
		}
		if contactInfo.ContactInfoType == "phone" && len(value) != 10 {
			return contactInfo, errors.New("phone number must be 10 digits")
		}
		contactInfo.ContactInfoValue = value
		return contactInfo, nil
	} else {
		return contactInfo, errors.New("invalid contact information value expected a string")
	}
}

func validateContactInfo(contactInfoType, contactInfoValue string) error {
	if contactInfoType == "" || contactInfoValue == "" {
		return errors.New("contact information type and Value cannot be empty")
	}
	if contactInfoType == "phone" && len(contactInfoValue) != 10 {
		return errors.New("phone number has to be 10 digits")
	}
	return nil
}

package students

import "time"

func isValidDate(dateOfBirth, monthOfBirth, yearOfBirth int) bool {
	if yearOfBirth > time.Now().Year() || yearOfBirth < 1900 {
		return false
	}
	if dateOfBirth > 31 || dateOfBirth <= 0 {
		return false
	}
	if monthOfBirth > 12 || monthOfBirth <= 0 {
		return false
	}
	if (monthOfBirth == 4 || monthOfBirth == 6 || monthOfBirth == 9 || monthOfBirth == 11) && dateOfBirth == 31 {
		return false
	}
	if monthOfBirth == 2 {
		if dateOfBirth >= 30 {
			return false
		}
		if dateOfBirth == 29 && !isLeapYear(yearOfBirth) {
			return false
		}
	}
	return true
}

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			return year%400 == 0
		}
		return true
	}
	return false
}

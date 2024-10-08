package students

import (
	"fmt"
	"time"
)

type dateOfBirth struct {
	date  int
	month int
	year  int
}
type Student struct {
	RollNo                  int
	FirstName               string
	LastName                string
	fullName                string
	dateOfBirth             dateOfBirth
	age                     int
	SemesterCGPA            []float64
	finalCGPA               float64
	semesterGrades          []string
	finalGrade              string
	YearOfEnrollment        int
	YearOfPassing           int
	numberOfYearsToGraduate int
}

var allStudents = []Student{}
var RollNos = 100

func NewStudent(
	firstName string,
	lastName string,
	dOB int,
	monthOfBirth int,
	yearOfBirth int,
	semCGPA []float64,
	yearOfEnrollment int,
	yearOfPassing int) *Student {

	if firstName == "" || lastName == "" {
		fmt.Println(("First name or last name cannot be empty"))
		return nil
	}

	if !isValidDate(dOB, monthOfBirth, yearOfBirth) {
		fmt.Println("Enter a valid date.")
		return nil
	}

	if !isValidCGPAArray(semCGPA) {
		return nil
	}

	if yearOfEnrollment < 1900 || yearOfEnrollment > time.Now().Year() {
		fmt.Println("Year of enrollment must be a valid year")
		return nil
	}

	if yearOfPassing < yearOfEnrollment {
		fmt.Println("Year of passing must be after year of enrollment")
		return nil
	}

	name := firstName + " " + lastName
	finalCGPA := GetFinalCGPA(semCGPA)
	var semesterGrades = []string{}
	GetSemesterGrades(semCGPA, &semesterGrades)
	ageCalc := time.Now().Year() - yearOfBirth
	finalGrade := GetGrade(finalCGPA)
	yearsToGraduation := yearOfPassing - time.Now().Year()
	if yearsToGraduation < 0 {
		yearsToGraduation = 0
	}

	stud := Student{
		RollNo:    RollNos,
		FirstName: firstName,
		LastName:  lastName,
		fullName:  name,
		dateOfBirth: dateOfBirth{
			date:  dOB,
			month: monthOfBirth,
			year:  yearOfBirth,
		},
		age:                     ageCalc,
		SemesterCGPA:            semCGPA,
		finalCGPA:               finalCGPA,
		semesterGrades:          semesterGrades,
		finalGrade:              finalGrade,
		YearOfEnrollment:        yearOfEnrollment,
		YearOfPassing:           yearOfPassing,
		numberOfYearsToGraduate: yearsToGraduation,
	}
	RollNos += 1
	fmt.Println("New Student Added!")
	allStudents = append(allStudents, stud)
	return &stud
}

func GetFinalCGPA(semCGPA []float64) float64 {
	sum := 0.0
	for _, val := range semCGPA {
		sum += val
	}
	return sum / float64(len(semCGPA))
}

func GetGrade(cgpa float64) string {
	if cgpa >= 9.5 {
		return "O"
	} else if cgpa >= 9.0 {
		return "A"
	} else if cgpa >= 8.0 {
		return "B"
	} else if cgpa >= 7.0 {
		return "C"
	} else if cgpa >= 6.0 {
		return "D"
	} else if cgpa >= 5.0 {
		return "E"
	} else {
		return "F"
	}
}

func GetSemesterGrades(semCGPA []float64, semesterGrades *[]string) {
	for _, cgpa := range semCGPA {
		*semesterGrades = append(*semesterGrades, GetGrade(cgpa))
	}
}

func isValidCGPAArray(semCGPA []float64) bool {
	if len(semCGPA) == 0 {
		fmt.Println("CGPA array cannot be empty")
		return false
	}

	for _, cgpa := range semCGPA {
		if cgpa < 0 || cgpa > 10 {
			fmt.Println("CGPA values must be between 0 and 10")
			return false
		}
	}
	return true
}

func GetAllStudents() {
	for _, student := range allStudents {
		fmt.Print("Student Roll No: ", student.RollNo, " ")
		fmt.Println(student)
	}
}

func GetStudentbyRollNo(rollNo int) {
	for _, student := range allStudents {
		if student.RollNo == rollNo {
			fmt.Println(student)
			return
		}
	}
	fmt.Println("No student with given roll number found!")
}

func DeleteStudentbyRollNo(rollNo int) {
	flag := 0
	allStudentsCopy := []Student{}
	for _, student := range allStudents {
		if student.RollNo != rollNo {
			allStudentsCopy = append(allStudentsCopy, student)
		} else {
			flag = 1
		}
	}
	allStudents = allStudentsCopy
	if flag == 1 {
		return
	}
	fmt.Println("No student with given roll number found!")
}

func (student *Student) UpdateStudentInfo(parameter int, newValue interface{}) {

	switch parameter {
	case 1:
		if value, ok := newValue.(string); ok {
			student.FirstName = value
			student.fullName = student.FirstName + " " + student.LastName

		} else {
			fmt.Println("Invalid type for First Name, expected a string.")
			return
		}
	case 2:
		if value, ok := newValue.(string); ok {
			student.LastName = value
			student.fullName = student.FirstName + " " + student.LastName

		} else {
			fmt.Println("Invalid type for Last Name, expected a string.")
			return
		}

	case 3:
		if value, ok := newValue.(int); ok {
			if isValidDate(value, student.dateOfBirth.month, student.dateOfBirth.year) {
				student.dateOfBirth.date = value

			} else {
				fmt.Println("Invalid Date modification. No such date exists.")
				return
			}
		} else {
			fmt.Println("Invalid type for Date of Birth, expected an integer.")
			return
		}

	case 4:
		if value, ok := newValue.(int); ok {
			if isValidDate(student.dateOfBirth.date, value, student.dateOfBirth.year) {
				student.dateOfBirth.month = value

			} else {
				fmt.Println("Invalid Date modification. No such date exists.")
				return
			}
		} else {
			fmt.Println("Invalid type for Month of Birth, expected an integer.")
			return
		}

	case 5:
		if value, ok := newValue.(int); ok {
			if isValidDate(student.dateOfBirth.date, student.dateOfBirth.month, value) {
				student.dateOfBirth.year = value
				student.age = time.Now().Year() - student.dateOfBirth.year

			} else {
				fmt.Println("Invalid Date modification. No such date exists.")
				return
			}
		} else {
			fmt.Println("Invalid type for Year of Birth, expected an integer.")
			return
		}

	case 6:
		if value, ok := newValue.(int); ok {
			if value >= 1900 || value <= time.Now().Year() {
				student.YearOfEnrollment = value

			} else {
				fmt.Println("Invalid Enrollment Year.")
				return
			}
		} else {
			fmt.Println("Invalid type for Year of Enrollment, expected an integer.")
			return
		}

	case 7:
		if value, ok := newValue.(int); ok {
			if student.YearOfPassing < value {
				student.YearOfPassing = value
				yearsToGraduation := student.YearOfPassing - time.Now().Year()
				if yearsToGraduation < 0 {
					yearsToGraduation = 0
				}
				student.numberOfYearsToGraduate = yearsToGraduation

			} else {
				fmt.Println("Invalid Graduation Year.")
				return
			}
		} else {
			fmt.Println("Invalid type for Year of Enrollment, expected an integer.")
			return
		}
	case 8:
		if value, ok := newValue.([]float64); ok {
			if isValidCGPAArray(value) {
				student.SemesterCGPA = value
				student.finalCGPA = GetFinalCGPA(student.SemesterCGPA)
				var semesterGrades = []string{}
				GetSemesterGrades(student.SemesterCGPA, &semesterGrades)
				student.finalGrade = GetGrade(student.finalCGPA)

			} else {
				fmt.Println("Invalid value for semester CGPA")
				return

			}
		}
	default:
		fmt.Println("Please enter a valid number from 1-8.")
	}
	// student.fullName = student.FirstName + " " + student.LastName
	// student.finalCGPA = GetFinalCGPA(student.SemesterCGPA)
	// var semesterGrades = []string{}
	// GetSemesterGrades(student.SemesterCGPA, &semesterGrades)
	// student.finalGrade = GetGrade(student.finalCGPA)
	// student.age = time.Now().Year() - student.dateOfBirth.year
	// student.numberOfYearsToGraduate = student.YearOfPassing - time.Now().Year()
	DeleteStudentbyRollNo(student.RollNo)
	allStudents = append(allStudents, *student)
}

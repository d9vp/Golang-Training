package main

import "dev/students"

func main() {
	student1 := students.NewStudent("Dev", "Patel", 9, 5, 2002, []float64{9.5, 9.0, 8.3, 8.5, 7.5, 9.0, 8.2, 6.8}, 2023, 2027)
	_ = students.NewStudent("Devnot", "NotPatel", 31, 10, 2000, []float64{9.2, 6.0, 7.3, 4.5, 7.5, 5.0, 8.2, 6.8}, 2023, 2027)
	_ = students.NewStudent("Krish", "Pandya", 29, 2, 2000, []float64{8.2, 8.0, 7.3, 9.5, 7.5, 6.2, 8.6, 7.8}, 2019, 2023)

	students.GetAllStudents()
	students.GetStudentbyRollNo(100)
	students.DeleteStudentbyRollNo(101)
	students.GetAllStudents()

	// updateMap := map[int]string{1: "First Name",
	// 	2: "Last Name",
	// 	3: "Date of Birth",
	// 	4: "Month of Birth",
	// 	5: "Year of Birth",
	// 	6: "Year of Enrollment",
	// 	7: "Year of Graduation",
	// 	8: "Semester CGPAs",
	// }

	student1.UpdateStudentInfo(3, 31)
	student1.UpdateStudentInfo(4, 11)
	students.GetAllStudents()

}

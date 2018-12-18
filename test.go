package main

import "fmt"
import "os"

type Student struct {
	name string
	age  int
}

func addStudentRecord(student map[int]Student, size int) {
	var tempStudent Student
	for i := 1; i <= size; i++ {
		fmt.Printf("Record No: %d\n", i)
		fmt.Print("Enter your name: ")
		fmt.Scan(&tempStudent.name)
		fmt.Print("\nEnter your age ")
		fmt.Scan(&tempStudent.age)
		student[i] = tempStudent
	}
}

func viewAllStudentRecord(student map[int]Student, size int) {
	for i := 1; i <= size; i++ {
		_, found := student[i]
		if found {
			fmt.Printf("\n%d: student  name is: %s and his age is: %d\n", i, student[i].name, student[i].age)
		} else {
			continue
		}
	}

}

func searchStudentRecord(student map[int]Student, size int) {
	var studentName string
	var flag int
	fmt.Printf("\nEnter the student Name you want to search : ")
	fmt.Scanln(&studentName)
	for i := 1; i <= size; i++ {
		if student[i].name == studentName {
			foundRecord := student[i]
			fmt.Printf("\nThe founded student record\n Student Name: %s\nStudent Age: %d", foundRecord.name, foundRecord.age)
			flag = 0
			break
		} else {
			flag = 1
		}
	}
	if flag == 1 {
		fmt.Print("\nRecord Not Found")
	}
}

func updateStudentRecord(student map[int]Student, size int) {
	var tempStudent Student
	var flag int
	fmt.Printf("\nEnter the student Name you want to search : ")
	fmt.Scanln(&tempStudent.name)
	for i := 1; i <= size; i++ {
		if student[i].name == tempStudent.name {
			fmt.Print("\nEnter the new name: ")
			fmt.Scan(&tempStudent.name)
			fmt.Print("\nEnter the new age ")
			fmt.Scan(&tempStudent.age)
			student[i] = tempStudent
			flag = 0
			break
		} else {
			flag = 1
		}
	}
	if flag == 1 {
		fmt.Print("\nRecord Not Found")
	} else {
		fmt.Print("\nYour record has been updated")
	}
}

func deleteStudentRecord(student map[int]Student, size int) {
	var tempStudent Student
	var flag int
	fmt.Printf("\nEnter the student Name you want to search : ")
	fmt.Scanln(&tempStudent.name)
	for i := 1; i <= size; i++ {
		if student[i].name == tempStudent.name {
			delete(student, i)
			flag = 0
			break
		} else {
			flag = 1
		}
	}
	if flag == 1 {
		fmt.Print("\nRecord Not Found")
	} else {
		fmt.Print("\nYour record has been deleted")
	}
}

func menu(student map[int]Student, size int) {
	var choice int
	fmt.Printf("\n\n\nFor Addition record press: 1")
	fmt.Printf("\nFor Search record press: 2")
	fmt.Printf("\nFor update record press: 3")
	fmt.Printf("\nFor View records press: 4")
	fmt.Printf("\nFor Delete record press: 5")
	fmt.Printf("\nFor exit Press any key except the choices of menu: ")
	fmt.Printf("\nEnter Your choice: ")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		addStudentRecord(student, size)
	case 2:
		fmt.Println("search")
		searchStudentRecord(student, size)
	case 3:
		fmt.Println("update")
		updateStudentRecord(student, size)
	case 4:
		fmt.Println("view")
		viewAllStudentRecord(student, size)
	case 5:
		deleteStudentRecord(student, size)
	default:
		fmt.Printf("default")
		os.Exit(0)
	}
}
func main() {
	var recordNum int
	fmt.Printf("Enter the number of records you want to enter: ")
	fmt.Scanln(&recordNum)
	student := make(map[int]Student)
	for {
		menu(student, recordNum)
	}
}


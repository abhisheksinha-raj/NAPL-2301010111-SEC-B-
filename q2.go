package main

import "fmt"


func modifyValue(num *int) {
	*num = *num + 10
}

// Structure
type Student struct {
	Name string
	Age  int
	marks float64
}

func main() {

	// 1. Referencing and Dereferencing
	num := 20

	fmt.Println("Original value:", num)
	fmt.Println("Address of num:", &num)
	fmt.Println("Value using pointer:", *(&num))

	// 2. Pass-by-reference using pointer
	fmt.Println("\nBefore function call:", num)

	modifyValue(&num)

	fmt.Println("After function call:", num)

	// 3. Allocate a struct using new()
	student := new(Student)

	// Access and modify fields through pointer
	student.Name = "Abhishek"
	student.Age = 21
	student.marks = 85.5
	fmt.Println("\nStudent Details:")
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("Marks:", student.marks)
}

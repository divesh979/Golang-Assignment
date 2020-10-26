package main

import (
	"fmt"
	"strconv"
)

func main() {
	//	Here we use the 2 different return values from the call with multiple assignment.
	empid := 1001
	empName, address, phoneNumber := returnMultipleValues(empid)
	fmt.Println("Returned Values of an Employee are: ", empName, address, phoneNumber)

	value := 15
	// Calling callByValue Method
	callByValue(value)
	fmt.Println("CallbyValue result: ", value) //Output will be "15"(i.e same Value as passed to Function) only as calling By Value

	// Calling callByReference Method
	callByReference(&value)
	fmt.Println("CallByReference result: ", value) //Output will be "15"(i.e same Value as passed to Function) only as calling By Value

	//FunctionClosures
	nextInteger := functionClosures()
	fmt.Println(nextInteger())
	fmt.Println(nextInteger())
	fmt.Println(nextInteger())
	newIntegers := functionClosures()
	fmt.Println(newIntegers())

	// Actual & Formal Parameters
	Sum := actualParameters()
	fmt.Println("Sum of Arguments:  ", Sum)

	//variable Declaration & Definition
	difference := variableDeclarationAndDefinition()
	fmt.Println("difference of Arguments:  ", difference)

	//Pointer to Poiter
	a, p, pp := pointerOnPointer()
	fmt.Println("value of a:", a)
	fmt.Println("address of a: ", &a)
	fmt.Println("value of p:", *p)
	fmt.Println("address of p: ", &p)

	// Dereferencing a pointer to pointer
	fmt.Println("*pp = ", *pp)
	fmt.Println("**pp = ", **pp)

	//map Example
	mapExample()

	//Int to Float Number
	var integerNumber int = 10
	floatNumber := intToFloat(integerNumber)
	fmt.Printf("float Conversion of integer Number = %.2f\n", floatNumber)

	var FloatNumber float64 = 10.11
	IntegerNumber := floatToInt(FloatNumber)
	fmt.Println("Integer Conversion of float Number ", IntegerNumber)

	stringValue := intToString(integerNumber)
	fmt.Println("stringValue of integer Number: ", stringValue)

	var StringValue string = "87"
	IntegerValue, err := stringToInt(StringValue)
	if err != nil {
		fmt.Println("Error Occured while Converting from String to int Value", err)
	} else {
		fmt.Println("IntegerValue of string Number: ", IntegerValue)
	}

}

//------------------ FirstQuestion ------------------//
//Return multiple values from a function
func returnMultipleValues(empid int) (string, string, int) {
	var empName string
	var address string
	var phoneNumber int

	if empid == 1001 {
		empName = "Aditya Ray Kapoor"
		address = "Mumbai"
		phoneNumber = 1234567890
	}
	return empName, address, phoneNumber
}

//------------------ SecondQuestion -------------------//
//Ways to pass parameters to a method
//Two ways to pass arguments to the function i.e. Call by Value and Call By Reference

// Go program to illustrate the concept of the callByValue
func callByValue(value int) {
	value = 70
}

// Go program to illustrate the concept of the callByReference
func callByReference(value *int) {
	*value = 70
}

//----------------Third Question----------------------//
//By default Golang Uses CallbyValue to pass the paramters to a
// method which is Demonstrated in above example

//----------------Fourth Question----------------------//
//Function closures
func functionClosures() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//----------------Fifth Question----------------------//
// The parameters passed to function are called Actual Parameters while
// the parameters received by the function are called Formal Parameters
// Here in Below formalParameters function a & b are formal Parameters
func formalParameters(a, b int) int {
	return a + b
}

func actualParameters() int {
	x, y := 10, 20
	// x & y are actual Arguments which are passed to formalParameters method
	sum := formalParameters(x, y)
	return sum
}

//----------------Sixth Question----------------------//
// Variable declaration and variable definition
func variableDeclarationAndDefinition() int {
	//Variable declaration in go
	var a int
	//variable definition in golang
	var b int = 10
	return b - a
}

//----------------Seventh Question----------------------//
//A pointer can point to a variable of any type. It can point to another pointer as well.
//The following example illustrates the same
func pointerOnPointer() (float64, *float64, **float64) {
	var a = 7.98
	var p = &a
	var pp = &p
	return a, p, pp
}

//----------------Eighth Question----------------------//
func mapExample() {
	var employees = map[int]string{
		1001: "John",
		1002: "Steve",
		1003: "Maria",
	}
	printEmployee(employees, 1001)
	printEmployee(employees, 1010)

	if isEmployeeExists(employees, 1002) {
		fmt.Println("EmployeeId 1002 found")
	}
}

func printEmployee(employees map[int]string, employeeId int) {
	if name, ok := employees[employeeId]; ok {
		fmt.Printf("name = %s, ok = %v\n", name, ok)
	} else {
		fmt.Printf("EmployeeId %d not found\n", employeeId)
	}
}

func isEmployeeExists(employees map[int]string, employeeId int) bool {
	_, ok := employees[employeeId]
	return ok
}

//----------------Ninth Question----------------------//
func intToFloat(integerNumber int) float64 {
	floatNumber := float64(integerNumber)
	return floatNumber
}

func floatToInt(floatNumber float64) int {
	integerNumber := int(floatNumber)
	return integerNumber
}

func intToString(integerNumber int) string {
	stringValue := strconv.Itoa(integerNumber)
	return stringValue
}

func stringToInt(stringValue string) (int, error) {
	integerNumber, err := strconv.Atoi(stringValue)
	if err != nil {
		return 0, err
	}
	return integerNumber, nil
}

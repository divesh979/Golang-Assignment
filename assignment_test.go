package main

import (
	"fmt"
	"testing"
)

func TestReturnMultipleValues(t *testing.T) {
	tests := []struct {
		name  string
		empid int
	}{
		{"Empid which exists", 1001},
		{"Empid doesn't exists", 1002},
	} //EndTestCases
	for _, testcase := range tests {
		empName, address, phoneNumber := returnMultipleValues(testcase.empid)
		if testcase.empid == 1001 {
			if empName != "Aditya Ray Kapoor" && address != "Mumbai" && phoneNumber != 1234567890 {
				t.Error()
			}
		}
	}
}

func TestCallByValue(t *testing.T) {
	value := 19
	callByValue(value)
	if value != 19 {
		t.Error()
	}
}

func TestCallByReference(t *testing.T) {
	value := 19
	callByReference(&value)
	if value != 70 {
		t.Error()
	}
}
func TestFunctionClosures(t *testing.T) {
	nextInteger := functionClosures()
	fmt.Println(nextInteger())
	newIntegers := functionClosures()
	fmt.Println(newIntegers())
}

func TestFormalParameters(t *testing.T) {
	a := 10
	b := 20
	value := formalParameters(a, b)
	if value != 30 {
		t.Error()
	}
}

func TestActualParameters(t *testing.T) {
	value := actualParameters()
	if value != 30 {
		t.Error()
	}
}

func TestVariableDeclarationAndDefinition(t *testing.T) {
	value := variableDeclarationAndDefinition()
	if value != 10 {
		t.Error()
	}
}

func TestMapExample(t *testing.T) {
	mapExample()
}

func TestPointerOnPointer(t *testing.T) {
	a, p, pp := pointerOnPointer()
	if a != 7.98 && *p != 7.98 && **pp != 7.98 {
		t.Error()
	}
}

func TestIsEmployeeExists(t *testing.T) {
	var employees = map[int]string{
		1001: "John",
		1002: "Steve",
		1003: "Maria",
	}
	printEmployee(employees, 1001)
	printEmployee(employees, 1004)
	result := isEmployeeExists(employees, 1001)
	if result == false {
		t.Error()
	}
	result = isEmployeeExists(employees, 1004)
	if result == true {
		t.Error()
	}

}

func TestIntToFloat(t *testing.T) {
	var integerNumber int = 10
	floatNumber := intToFloat(integerNumber)
	fmt.Println("floatNumber", floatNumber)
	if floatNumber != 10.00 {
		t.Error()
	}
}

func TestFloatToInt(t *testing.T) {
	var floatNumber float64 = 10.11
	IntegerNumber := floatToInt(floatNumber)
	if IntegerNumber != 10 {
		t.Error()
	}
}

func TestIntToString(t *testing.T) {
	var integerNumber int = 10
	stringValue := intToString(integerNumber)
	if stringValue != "10" {
		t.Error()
	}
}

func TestStringToInt(t *testing.T) {
	var stringValue string = "76"
	IntegerNumber, err := stringToInt(stringValue)
	if IntegerNumber != 76 && err != nil {
		t.Error()
	}
}

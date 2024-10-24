package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Adder interface {
	Add(a, b int) int
}

type Subtractor interface {
	Subtract(a, b int) int
}

type DefaultAdder struct{}

func (d DefaultAdder) Add(a, b int) int {
	return a + b
}

type DefaultSubtractor struct{}

func (d DefaultSubtractor) Subtract(a, b int) int {
	return a - b
}

func DynamicLoadAdder() Adder {
	return DefaultAdder{}
}

func DynamicLoadSubtractor() Subtractor {
	return DefaultSubtractor{}
}

func ParseExpression(expr string) (int, int, string, error) {

	expr = strings.ReplaceAll(expr, " ", "")

	// find the operator
	// split the string into two numbers based on the operator
	var operator string
	if strings.Contains(expr, "+") {
		operator = "+"
	} else if strings.Contains(expr, "-") {
		operator = "-"
	} else {
		return 0, 0, "", fmt.Errorf("unsupported")
	}

	parts := strings.Split(expr, operator)
	if len(parts) != 2 {
		return 0, 0, "", fmt.Errorf("invalid expression")
	}

	// Convert the strings to integers
	a, err1 := strconv.Atoi(parts[0])
	b, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return 0, 0, "", fmt.Errorf("invalid numbers")
	}

	return a, b, operator, nil
}

func main() {
	// Input the mathematical expression
	var expr string
	fmt.Println("Enter a mathematical expression (e.g., '10 + 5' or '10 - 5'):")
	fmt.Scanln(&expr)

	// Parse the expression to extract the numbers and operator
	a, b, operator, err := ParseExpression(expr)
	if err != nil {
		fmt.Println("Error parsing expression:", err)
		return
	}

	if operator == "+" {
		adder := DynamicLoadAdder()
		result := adder.Add(a, b)
		fmt.Printf("The result of %d + %d is: %d\n", a, b, result)
	} else if operator == "-" {
		subtractor := DynamicLoadSubtractor()
		result := subtractor.Subtract(a, b)
		fmt.Printf("The result of %d - %d is: %d\n", a, b, result)
	}
}

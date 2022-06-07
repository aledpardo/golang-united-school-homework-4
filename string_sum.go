package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.Trim(input, " ")
	var numbers [2]int
	var index int
	var negative bool
	for _, v := range input {
		s := string(v)
		if s == " " {
			continue
		}
		if s == "-" {
			negative = true
			continue
		}
		if s == "+" {
			continue
		}
		n, e := strconv.ParseInt(s, 10, 32)
		if e != nil {
			return "", e // TODO create Error
		}
		if index > 1 {
			return "", nil // TODO create Error
		}
		if negative {
			n = n * -1
			negative = false
		}
		numbers[index] = int(n)
		index++
	}
	if index == 0 {
		return "", errorEmptyInput
	}
	if index == 1 || index > 2 {
		return "", errorNotTwoOperands
	}
	output = fmt.Sprint(numbers[0] + numbers[1])
	return output, nil
}

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
	input = strings.ReplaceAll(input, " ", "")
	
	if len(input) == 0 {
		err = fmt.Errorf("input is not valid: %w", errorEmptyInput)
		return
	}
	
	strNums := Split(
		input, []rune{'+', '-'})
	
	var nums []int
	for _, s := range strNums {
		n, ok := strconv.Atoi(s)
		if ok != nil {
			err = fmt.Errorf("inappropriate symbols: %w", ok)
			return
		}
		nums = append(nums, n)
	}
	
	if len(strNums) != 2 {
		err = fmt.Errorf("invalid input: %w", errorNotTwoOperands)
		return
	}
	
	var sign = 1
	var numToAdd int8 = 0
	var sum int
	
	for i := 0; i < len(input); i++ {
		if input[i] == '-' {
			sign *= -1
		} else if input[i] == '+' {
			continue
		} else {
			i += len(strNums[numToAdd]) - 1
			sum += nums[numToAdd] * sign
			numToAdd++
			sign = 1
		}
	}
	output = strconv.Itoa(sum)
	return
}

func Split(str string, separators []rune) []string {
	f := func(r rune) bool {
		for _, s := range separators {
			if r == s {
				return true
			}
		}
		return false
	}
	return strings.FieldsFunc(str, f)
}

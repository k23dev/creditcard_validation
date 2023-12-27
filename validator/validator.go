package validator

import (
	"strconv"
	"unicode/utf8"
)

// algorithm to check the credit card
//https://en.wikipedia.org/wiki/Luhn_algorithm]

func Validate(ccNumber int) bool {

	//ccNumberLength = 16

	ccNumberStr := strconv.Itoa(ccNumber)

	sum := 0

	var lastDigit int

	for key, digit := range ccNumberStr {

		if key == len(ccNumberStr)-1 {
			lastDigit = convertToInt(digit)
			break
		}
		num := convertToInt(digit)
		if key%2 == 0 {
			if num > 4 {
				// This replaces the sumatory of the result of digit * 2 and then decomposing and sum those numbers.
				// example 7*2= 14 -> 1+4 = 5 | 7*2=14 -> 14-9 = 5
				sum = sum + 2*num - 9
			} else {
				sum = sum + 2*num
			}
		} else {
			sum = sum + num
		}
	}

	return (lastDigit == 10-(sum%10))

}

func convertToInt(numRune rune) int {

	buf := make([]byte, 1)
	_ = utf8.EncodeRune(buf, numRune)
	num, _ := strconv.Atoi(string(buf))

	return num

}

package main

import (
	"fmt"
	"strconv"
)

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}

func addStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	carry := 0
	result := ""

	for i >= 0 || j >= 0 {
		sum := carry
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}

		result += strconv.Itoa(sum % 10)
		carry = sum / 10
	}

	if carry != 0 {
		result += strconv.Itoa(carry)
	}

	return reverseString(result)
}

func main() {
	fmt.Println(addStrings("94", "94"))
}

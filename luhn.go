package main

import (
	"fmt"
	"strconv"
)

func validCard(cardNum string) bool {
	if len(cardNum) > 19 || len(cardNum) == 0 {
		return false
	}
	sum := 0
	modifyDigit := false
	for i := len(cardNum) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(cardNum[i]))
		if err != nil {
			fmt.Printf("Error converting card digit '%s' to int: %v\n", string(cardNum[i]), err)
			return false
		}
		if modifyDigit {
			digit = digit * 2
			if digit > 9 {
				digit = digit - 9
			}
		}
		sum = sum + digit
		modifyDigit = !modifyDigit
	}
	if sum%10 == 0 {
		return true
	}
	return false
}

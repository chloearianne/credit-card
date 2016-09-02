package main

import "strconv"

//  Account is a representation of one person's credit card account.
// Error should only ever be set to 'true' if the card number is invalid
// (upon attempting to create an account).
type Account struct {
	Balance int
	Limit   int
	Error   bool
}

// Charge adds an amount to the account's balance if it won't exceed the limit.
// It does nothing if the limit would be exceeded by adding the amount.
func (a *Account) Charge(amount string) {
	if a.Error {
		return
	}
	newBalance := a.Balance + moneyStringToInt(amount)
	if newBalance > a.Limit {
		return
	}
	a.Balance = newBalance
}

// Credit subtracts an amount from the account's balance.
func (a *Account) Credit(amount string) {
	if a.Error {
		return
	}
	a.Balance -= moneyStringToInt(amount)
}

// createAccount takes in a card number and a beginning balance and creates a new
// Account with that balance if the card number is valid. If the card isn't valid,
// it returns an account with an empty balance with Error set to true.
func createAccount(cardNum, limit string) *Account {
	acc := Account{}
	if !validCard(cardNum) {
		acc.Error = true
		return &acc
	}
	num := moneyStringToInt(limit)
	acc.Limit = num
	acc.Balance = 0
	acc.Error = false
	return &acc
}

// moneyStringToInt takes in a string amount and returns the integer equivalent.
// It assumes amount is well formed, i.e. it begins with a $, is non negative, and
// represents a valid integer.
func moneyStringToInt(amount string) int {
	amount = amount[1:]            // strip $ sign off of beginning
	num, _ := strconv.Atoi(amount) // ignoring error on assumption of valid input
	return num
}

package main

import "testing"

func TestValidAdd(t *testing.T) {
	acc := createAccount("4111111111111111", "$400")
	if acc.Error {
		t.Errorf("Error was set on valid card")
	}
	if acc.Balance != 0 {
		t.Errorf("Expected balance to be 0, but got %d", acc.Balance)
	}
	if acc.Limit != 400 {
		t.Errorf("Expected limit to be 400, but got %d", acc.Balance)
	}
}

func TestInvalidAdd(t *testing.T) {
	acc := createAccount("badcardnumber", "$600")
	if !acc.Error {
		t.Errorf("Error wasn't set on invalid card")
	}
	if acc.Limit != 0 {
		t.Errorf("Expected Limit to be 0, but got set to %d", acc.Limit)
	}
}

func TestValidCharges(t *testing.T) {
	acc := Account{
		Limit:   1000,
		Balance: 400,
		Error:   false,
	}
	acc.Charge("$250")
	if acc.Balance != 650 {
		t.Fatalf("After charging 250, expected sum of 650, but got %d", acc.Balance)
	}
	acc.Charge("$1")
	if acc.Balance != 651 {
		t.Errorf("After charging 1, expected sum of 651, but got %d", acc.Balance)
	}
}

func TestInvalidChargeOverLimit(t *testing.T) {
	acc := Account{
		Limit:   450,
		Balance: 375,
		Error:   false,
	}
	acc.Charge("$25")
	if acc.Balance != 400 {
		t.Fatalf("After charging 25, expected sum of 400, but got %d", acc.Balance)
	}
	acc.Charge("$250")
	if acc.Balance != 400 {
		t.Fatalf("Expected balance not to change, but got %d", acc.Balance)
	}
	acc.Charge("$1")
	if acc.Balance != 401 {
		t.Errorf("After charging 1, expected sum of 401, but got %d", acc.Balance)
	}
}

func TestInvalidChargeBadAccount(t *testing.T) {
	acc := Account{
		Limit:   0,
		Balance: 0,
		Error:   true,
	}
	acc.Charge("$25")
	if acc.Balance != 0 {
		t.Fatalf("Expected balance to be unchanged, but got %d", acc.Balance)
	}
}

func TestValidCredit(t *testing.T) {
	acc := Account{
		Limit:   1000,
		Balance: 400,
		Error:   false,
	}
	acc.Credit("$250")
	if acc.Balance != 150 {
		t.Fatalf("After crediting 250, expected sum of 150, but got %d", acc.Balance)
	}
	acc.Credit("$1")
	if acc.Balance != 149 {
		t.Errorf("After crediting 1, expected sum of 149, but got %d", acc.Balance)
	}
}

func TestInvalidCredit(t *testing.T) {
	acc := Account{
		Limit:   0,
		Balance: 0,
		Error:   true,
	}
	acc.Credit("$250")
	if acc.Balance != 0 {
		t.Fatalf("Expected balance to be unchanged, but got %d", acc.Balance)
	}
}

func TestMultipleTransactions(t *testing.T) {
	acc := createAccount("5454545454545454", "$550")
	if acc.Error {
		t.Fatalf("Error was set on valid card")
	}
	if acc.Balance != 0 {
		t.Fatalf("Expected balance to be 0, but got %d", acc.Balance)
	}
	if acc.Limit != 550 {
		t.Fatalf("Expected limit to be 400, but got %d", acc.Balance)
	}
	acc.Credit("$250")
	if acc.Balance != -250 {
		t.Fatalf("Expected balance to be -250, but got %d", acc.Balance)
	}
	acc.Charge("$1000")
	if acc.Balance != -250 {
		t.Fatalf("Expected balance to be -250, but got %d", acc.Balance)
	}
	acc.Charge("$500")
	if acc.Balance != 250 {
		t.Fatalf("Expected balance to be 250, but got %d", acc.Balance)
	}
	acc.Credit("$15")
	if acc.Balance != 235 {
		t.Fatalf("Expected balance to be 235, but got %d", acc.Balance)
	}
	acc.Charge("$50")
	if acc.Balance != 285 {
		t.Fatalf("Expected balance to be 235, but got %d", acc.Balance)
	}
	acc.Charge("$300")
	if acc.Balance != 285 {
		t.Fatalf("Expected balance to be 235, but got %d", acc.Balance)
	}
}

func TestMultipleAccounts(t *testing.T) {
	acc1 := createAccount("5454545454545454", "$550")
	if acc1.Error {
		t.Fatalf("Error was set on valid card")
	}
	if acc1.Balance != 0 {
		t.Fatalf("Expected balance to be 0, but got %d", acc1.Balance)
	}
	if acc1.Limit != 550 {
		t.Fatalf("Expected limit to be 400, but got %d", acc1.Limit)
	}
	acc2 := createAccount("4111111111111111", "$1000")
	if acc2.Error {
		t.Fatalf("Error was set on valid card")
	}
	if acc2.Balance != 0 {
		t.Fatalf("Expected balance to be 0, but got %d", acc2.Balance)
	}
	if acc2.Limit != 1000 {
		t.Fatalf("Expected limit to be 400, but got %d", acc2.Limit)
	}

	acc1.Charge("$250")
	if acc1.Balance != 250 {
		t.Fatalf("Expected balance to be 250, but got %d", acc1.Balance)
	}
	acc2.Charge("$120")
	if acc2.Balance != 120 {
		t.Fatalf("Expected balance to be 120, but got %d", acc2.Balance)
	}
	acc1.Credit("$50")
	if acc1.Balance != 200 {
		t.Fatalf("Expected balance to be 200, but got %d", acc1.Balance)
	}
	acc2.Credit("$100")
	if acc2.Balance != 20 {
		t.Fatalf("Expected balance to be 20, but got %d", acc2.Balance)
	}
}

func TestMoneyStringToInt(t *testing.T) {
	num := moneyStringToInt("$234")
	if num != 234 {
		t.Errorf("Expected 234, got %d", num)
	}
}

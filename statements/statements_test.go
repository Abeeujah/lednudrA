package statements_test

import (
	"statements"
	"testing"
)

func TestWithdraw(t *testing.T) {
	t.Parallel()

	balance := 100
	want := 80

	got, err := statements.Withdraw(balance, 20)

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("Expected %d, got %d", want, got)
	}
}

func TestWithdrawInvalid(t *testing.T) {
	t.Parallel()

	balance := 20
	amount := 100

	_, err := statements.Withdraw(balance, amount)

	if err == nil {
		t.Fatal("Expected an error for Overdraw, got nil.")
	}
}

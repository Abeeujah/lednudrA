package creditcard_test

import (
	"creditcard"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	want := "123456789"

	cc, err := creditcard.New(want)
	got := cc.Number()

	if err != nil {
		t.Fatal(err)
	}

	if want != got {
		t.Errorf("Card values don't match, expected %q, got %q", want, got)
	}
}

func TestNewInvalid(t *testing.T) {
	t.Parallel()

	_, err := creditcard.New("")

	if err == nil {
		t.Fatal("Expected an error for Invalid credit card details, got nil.")
	}
}

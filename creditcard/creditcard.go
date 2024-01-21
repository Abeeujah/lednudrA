package creditcard

import (
	"errors"
)

type card struct {
	number string
}

func New(number string) (card, error) {
	if number == "" {
		return card{}, errors.New("Error, Credit card details cannot be empty")
	}
	return card{number}, nil
}

func (cc card) Number() string {
	return cc.number
}

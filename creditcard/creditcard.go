package creditcard

import "errors"

// private struct and hence cannot be
// initialized directly from outside of this package
type card struct {
	number string
}

func New(number string) (card, error) {
	if number == "" {
		return card{}, errors.New("number cannot be empty")
	}
	return card{number}, nil
}

func (c *card) Number() string {
	return c.number
}

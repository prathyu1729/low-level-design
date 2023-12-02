package atm

import "errors"

type idle struct {
	atm *Atm
}

func (i *idle) authenticate(card Card, pin int) error {
	if card.pin != pin {
		return errors.New("Invalid Pin")
	}

	i.atm.setState(i.atm.authenticated)
	return nil
}

func (i *idle) withdraw(amount int64) error {
	return errors.New("Enter your card first")
}

func (i *idle) checkBalance() (float64, error) {
	return 0, errors.New("Enter your card first")
}

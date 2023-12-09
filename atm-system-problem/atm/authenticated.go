package atm

import (
	"errors"
)

type authenticated struct {
	atm  *Atm
	card *Card
}

func (a *authenticated) setCard(card *Card) {
	a.card = card
}

func (a *authenticated) unSetCard() {
	a.card = nil
}

func (a *authenticated) authenticate(pin string) error {
	return errors.New("ATM already in use")
}

func (a *authenticated) checkBalance() (float64, error) {
	account := a.atm.cardMapping[a.card.cardNumber]
	return account.getBalance(), nil
}

func (a *authenticated) withdraw(amount float64) error {
	if amount == 0 {
		return errors.New("Invalid amount entered")
	}
	account := a.atm.cardMapping[a.card.cardNumber]
	if account.getBalance() < amount {
		return errors.New("Insufficient Balance")
	}

	if a.atm.atmBalance < amount {
		return errors.New("Request not supported at the moment")
	}

	flag, denominations := a.atm.amountChecker.getDenominations(amount)
	if !flag {
		return errors.New("Unsupported amount")
	}

	if err := a.atm.reduceBalance(amount); err != nil {
		return err
	}

	if err := a.atm.updateDenominations(denominations); err != nil {
		return err
	}

	a.unSetCard()
	a.atm.setState(a.atm.idle)
	return nil
}



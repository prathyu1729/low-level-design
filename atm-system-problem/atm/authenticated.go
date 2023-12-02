package atm

import "errors"

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

	//TODO: denomination check
	a.atm.atmBalance = a.atm.atmBalance - amount
	//TODO: update denomination
	a.unSetCard()
	a.atm.setState(a.atm.idle)
	return nil
}

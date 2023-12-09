package atm

import "sync"

type Atm struct {
	cardMapping   map[string]bankAccount
	atmBalance    float64
	denominations map[float64]int
	amountChecker *amountChecker
	idle          state
	authenticated state
	currentState  state
}

var atm Atm
var once sync.Once

func NewAtm() *Atm {
	once.Do(func() {
		atm = Atm{}
	})

	return &atm
}

func (a *Atm) Authenticate(card Card, pin int) error {
	return a.currentState.authenticate(card, pin)
}

func (a *Atm) Withdraw(amount float64) error {
	return a.currentState.withdraw(amount)
}

func (a *Atm) CheckBalance() (float64, error) {
	return a.currentState.checkBalance()
}

func (a *Atm) setState(state state) {
	a.currentState = state
}

func (a *Atm) getDenominations() []int {
	//TODO:
}

func (a *Atm) updateDenominations(denom map[float64]int) error {

}

func (a *Atm) reduceBalance(amount float64) error {

}

package atm

type state interface {
	authenticate(card Card, pin int) error
	withdraw(amount float64) error
	checkBalance() (float64, error)
}

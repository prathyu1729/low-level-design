package atm

type bankAccount interface {
	getBalance() float64
}

type savingsAccount struct {
	balance float64
}

type currentAccount struct {
	balance float64
}

func (s *savingsAccount) getBalance() float64 {
	return s.balance
}

func (c *currentAccount) getBalance() float64 {
	return c.balance
}

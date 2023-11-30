package parking_lot

import "fmt"

type CardPayment struct {
}

func (c CardPayment) MakePayment(amount float64) {
	fmt.Println("Paid %f with card", amount)
}

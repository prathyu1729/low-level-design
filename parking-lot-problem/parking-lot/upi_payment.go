package parking_lot

import "fmt"

type UpiPayment struct {
}

func (u UpiPayment) MakePayment(amount float64) {
	fmt.Println("Paid %f with upi", amount)
}

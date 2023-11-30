package parking_lot

import "time"

type Exit struct {
	Id int
}

const ParkingRate = 10

func (e Exit) AcceptTicket(ticket Ticket, payment Payment) {
	timeInSeconds := time.Since(ticket.StartTime).Seconds()
	amount := timeInSeconds * 10
	payment.MakePayment(amount)
}

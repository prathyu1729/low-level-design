package main

import (
	"fmt"
	parking_lot "parking-lot/parking-lot"
)

func main() {
	parkingLot := parking_lot.NewParkingLot()
	carA := parking_lot.CarVehicle{}
	ticket, err := parkingLot.GenerateTicket(parking_lot.EntryA, carA)
	if err != nil {
		fmt.Println(err.Error())
	}
	paymentA := parking_lot.CardPayment{}
	err = parkingLot.LeaveParking(ticket, parking_lot.ExitA, paymentA)
	if err != nil {
		fmt.Println(err.Error())
	}
}

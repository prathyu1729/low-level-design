package parking_lot

import "time"

type Entry struct {
	Id int
}

func (e Entry) GetTicket(slot ParkingSlot) Ticket {
	return Ticket{
		SlotId:    slot.getId(),
		StartTime: time.Now(),
	}
}

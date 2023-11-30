package parking_lot

import (
	"errors"
	"sync"
)

type ParkingLot struct {
	entries map[EntryPoint]Entry
	exits   map[ExitPoint]Exit
	slots   []ParkingSlot
}

var parkingLot *ParkingLot
var once sync.Once

func NewParkingLot() *ParkingLot {
	once.Do(func() {
		parkingLot = &ParkingLot{}
	})
	return parkingLot
}

func (p *ParkingLot) GenerateTicket(entryName EntryPoint, vehicle Vehicle) (Ticket, error) {
	entry, ok := p.entries[entryName]
	if !ok {
		return Ticket{}, errors.New("Invalid entry point")
	}
	slot, err := p.findSlot(vehicle)
	if err != nil {
		return Ticket{}, err
	}
	ticket := entry.GetTicket(slot)
	//TODO: update to not free
	return ticket, nil
}

func (p *ParkingLot) LeaveParking(ticket Ticket, exitPoint ExitPoint, payment Payment) error {
	exit, ok := p.exits[exitPoint]
	if !ok {
		return errors.New("Invalid exit point")
	}
	exit.AcceptTicket(ticket, payment)
	//TODO: mark slot as free
	return nil
}

func (p *ParkingLot) findSlot(vehicle Vehicle) (ParkingSlot, error) {
	for _, slot := range p.slots {
		if slot.isFree() && (slot.getType() == vehicle.getType()) {
			return slot, nil
		}
	}

	return nil, errors.New("Slot not available")
}

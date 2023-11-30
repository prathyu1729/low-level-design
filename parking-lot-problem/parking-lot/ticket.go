package parking_lot

import "time"

type Ticket struct {
	StartTime time.Time
	SlotId    int
}

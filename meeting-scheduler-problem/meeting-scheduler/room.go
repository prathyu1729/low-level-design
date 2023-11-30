package meeting_scheduler

type room struct {
	name            string
	capacity        int
	bookedIntervals []Interval
}

func newRoom(name string, capacity int) room {
	return room{
		name:            name,
		capacity:        capacity,
		bookedIntervals: []Interval{},
	}
}

func (r room) isIntervalFree(interval Interval) bool {
	start := interval.startTime.Unix()
	end := interval.endTime.Unix()
	for _, i := range r.bookedIntervals {
		bookedStart := i.startTime.Unix()
		bookedEnd := i.endTime.Unix()
		if (start >= bookedStart) && (start <= bookedEnd) || (end >= bookedStart && end <= bookedEnd) {
			return false
		}
	}
	return true
}

func (r room) canAccomodate(count int) bool {
	return count <= r.capacity
}

func (r room) addInterval(interval Interval) room {
	r.bookedIntervals = append(r.bookedIntervals, interval)
	return r
}

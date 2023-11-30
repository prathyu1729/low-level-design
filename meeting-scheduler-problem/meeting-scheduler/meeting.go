package meeting_scheduler

import (
	"fmt"
	"time"
)

type Interval struct {
	startTime time.Time
	endTime   time.Time
}

func NewInterval(startTime time.Time, endTime time.Time) Interval {
	return Interval{
		startTime: startTime,
		endTime:   endTime,
	}
}

type Meeting struct {
	interval     Interval
	participants []User
	room         room
	organizer    User
}

func (m Meeting) ShowDetails() {
	fmt.Println("********************")
	fmt.Println("Meeting Room: ", m.room.name)
	fmt.Println("Starting Time: ", m.interval.startTime)
	fmt.Println("End Time: ", m.interval.endTime)
	fmt.Println("Organizer: ", m.organizer.name)
	fmt.Println("********************")
}

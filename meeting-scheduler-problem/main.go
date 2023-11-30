package main

import (
	"fmt"
	meeting_scheduler "meeting-scheduler/meeting-scheduler"
	"time"
)

func main() {

	ms := meeting_scheduler.NewMeetingScheduler()
	ms.AddRoom("A", 5).AddRoom("B", 1)

	user1 := meeting_scheduler.NewUser("Prathyush")
	user2 := meeting_scheduler.NewUser("Aks")
	user3 := meeting_scheduler.NewUser("Anand")
	user4 := meeting_scheduler.NewUser("Abhi")
	user5 := meeting_scheduler.NewUser("Nandu")

	interval1 := meeting_scheduler.NewInterval(time.Date(2023, 12, 1, 1, 0, 0, 0, time.UTC),
		time.Date(2023, 12, 1, 2, 0, 0, 0, time.UTC))

	meeting1, err := ms.ScheduleMeeting(user1, interval1, []meeting_scheduler.User{user2, user3, user4, user5})
	if err != nil {
		fmt.Println(err.Error())
	}
	if meeting1 != nil {
		meeting1.ShowDetails()
	}
	meeting2, err := ms.ScheduleMeeting(user1, interval1, []meeting_scheduler.User{user2, user3, user4, user5})
	if err != nil {
		fmt.Println(err.Error())
	}
	if meeting2 != nil {
		meeting2.ShowDetails()
	}
}

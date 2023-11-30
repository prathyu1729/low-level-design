package meeting_scheduler

import (
	"errors"
	"sync"
)

type MeetingScheduler struct {
	meetings []Meeting
	rooms    map[string]room
	users    []User
}

var meetingScheduler MeetingScheduler
var once sync.Once

func NewMeetingScheduler() *MeetingScheduler {
	once.Do(func() {
		meetingScheduler = MeetingScheduler{
			meetings: []Meeting{},
			rooms:    map[string]room{},
			users:    []User{},
		}
	})

	return &meetingScheduler
}

func (m *MeetingScheduler) AddRoom(name string, capacity int) *MeetingScheduler {
	m.rooms[name] = newRoom(name, capacity)
	return m
}

func (m *MeetingScheduler) ScheduleMeeting(oragniser User, interval Interval, participants []User) (*Meeting, error) {
	roomName, err := m.searchRoom(interval, len(participants)+1)
	if err != nil {
		return nil, err
	}

	m.rooms[roomName] = m.rooms[roomName].addInterval(interval)
	newMeeting := Meeting{
		interval:     interval,
		participants: participants,
		organizer:    oragniser,
		room:         m.rooms[roomName],
	}
	m.meetings = append(m.meetings, newMeeting)
	return &newMeeting, nil
}

func (m *MeetingScheduler) searchRoom(interval Interval, capacity int) (string, error) {
	for _, room := range m.rooms {
		if room.isIntervalFree(interval) && room.canAccomodate(capacity) {
			return room.name, nil
		}
	}
	return "", errors.New("Room Unavailable")
}

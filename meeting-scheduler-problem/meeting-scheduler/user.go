package meeting_scheduler

type User struct {
	id   int
	name string
}

func NewUser(name string) User {
	return User{name: name}
}

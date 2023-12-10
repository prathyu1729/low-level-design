package pubsub

type Message struct {
	payload string
}

func NewMessage(payload string) Message {
	return Message{payload: payload}
}

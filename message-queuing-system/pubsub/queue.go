package pubsub

type queue struct {
	messages []Message
}

func newQueue() queue {
	return queue{messages: []Message{}}
}

func (q queue) enQueue(message Message) {
	q.messages = append(q.messages, message)
}

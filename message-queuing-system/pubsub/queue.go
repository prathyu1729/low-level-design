package pubsub

import "fmt"

type queue struct {
	messages []Message
}

func newQueue() *queue {
	return &queue{messages: []Message{}}
}

func (q *queue) enQueue(message Message) {
	fmt.Println("Queued message: ", message.payload)
	q.messages = append(q.messages, message)
}

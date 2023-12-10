package pubsub

import (
	"errors"
	"fmt"
	"sync"
)

type Publisher struct {
	observers map[observer]bool
	queue     *queue
	mu        sync.Mutex
}

func NewPublisher() *Publisher {
	return &Publisher{
		observers: map[observer]bool{},
		queue:     newQueue(),
	}
}

func (p *Publisher) Enqueue(message Message) {
	p.queue.enQueue(message)
}

func (p *Publisher) Subscribe(observer observer) {
	p.observers[observer] = true
}

func (p *Publisher) UnSubscribe(observer observer) error {
	if _, ok := p.observers[observer]; !ok {
		return errors.New("observer not present")
	}
	delete(p.observers, observer)
	return nil
}

func (p *Publisher) Notify() {
	fmt.Println("started notifying")
	messages := p.queue.messages
	fmt.Println(" no of messages", len(messages))
	for _, message := range messages {
		for observer, _ := range p.observers {
			observer.getMessage(message)
		}
	}
}

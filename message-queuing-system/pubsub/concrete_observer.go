package pubsub

import "fmt"

type ConcreteObserver struct {
	name         string
	nextObserver observer
}

func NewConcreteObserver(name string) *ConcreteObserver {
	return &ConcreteObserver{
		name: name,
	}
}

func (c *ConcreteObserver) SetNext(observer observer) {
	c.nextObserver = observer
}

func (c *ConcreteObserver) getMessage(message Message) {
	fmt.Println("Observer", c.name, "received", message.payload)
	if c.nextObserver != nil {
		c.nextObserver.getMessage(message)
	}
}

func (c *ConcreteObserver) getNext() observer {
	return c.nextObserver
}

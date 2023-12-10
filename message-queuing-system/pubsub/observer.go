package pubsub

type observer interface {
	SetNext(observer observer)
	getMessage(message Message)
	getNext() observer
}

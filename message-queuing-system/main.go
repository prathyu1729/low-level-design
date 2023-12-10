package main

import "queue/pubsub"

func main() {
	publisher := pubsub.NewPublisher()
	observer1 := pubsub.NewConcreteObserver("service1")
	observer2 := pubsub.NewConcreteObserver("service2")
	observer3 := pubsub.NewConcreteObserver("service3")
	observer4 := pubsub.NewConcreteObserver("service4")

	observer3.SetNext(observer4)

	//subscribe the observers
	publisher.Subscribe(observer1)
	publisher.Subscribe(observer2)
	publisher.Subscribe(observer3)

	//enqueue messages
	publisher.Enqueue(pubsub.NewMessage("payment  1 done"))
	publisher.Enqueue(pubsub.NewMessage("payment 5 done"))
	publisher.Enqueue(pubsub.NewMessage("payment 67 done"))

	//notify messages
	publisher.Notify()
}

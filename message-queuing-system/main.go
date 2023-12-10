package main

import (
	"fmt"
	"queue/pubsub"
)

func main() {
	publisher := pubsub.NewPublisher()
	observer1 := pubsub.NewConcreteObserver("service1")
	observer2 := pubsub.NewConcreteObserver("service2")
	observer3 := pubsub.NewConcreteObserver("service3")
	observer4 := pubsub.NewConcreteObserver("service4")
	observer5 := pubsub.NewConcreteObserver("service5")

	observer4.SetNext(observer3)
	observer3.SetNext(observer2)
	observer2.SetNext(observer1)

	//subscribe the observers
	publisher.Subscribe(observer4)
	publisher.Subscribe(observer5)

	//enqueue messages
	publisher.Enqueue(pubsub.NewMessage("payment  1 done"))
	publisher.Enqueue(pubsub.NewMessage("payment 5 done"))
	publisher.Enqueue(pubsub.NewMessage("payment 67 done"))

	if err := publisher.UnSubscribe(observer5); err != nil {
		fmt.Println(err.Error())
	}

	//notify messages
	publisher.Notify()
}

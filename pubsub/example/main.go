package main

import (
	"fmt"
	"strconv"

	"github.com/design-patterns/pubsub"
)

func main() {

	subscribers := make([]*pubsub.Subscriber, 5)
	for i := 1; i <= 5; i++ {
		subscribers[i-1] = pubsub.NewSubscriberWithFunc(handler)
	}

	broker := pubsub.NewBroker()
	for i, sub := range subscribers {
		broker.Subscribe(sub, "topic:"+strconv.Itoa(i+1))
	}

	for i := 1; i <= 5; i++ {
		broker.Publish("topic:"+strconv.Itoa(i), "Europe")
	}
}

func handler(msg string) {
	fmt.Printf("Hello %s\n", msg)
}

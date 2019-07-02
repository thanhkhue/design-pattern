package pubsub

import "sync"

// Broker distributes a message to subscribers who listen to topic
type Broker struct {
	mu          *sync.RWMutex
	publisher   *Publisher
	subscribers map[string][]*Subscriber
}

// NewBroker new message broker
func NewBroker() *Broker {
	return &Broker{
		mu:          new(sync.RWMutex),
		publisher:   new(Publisher),
		subscribers: make(map[string][]*Subscriber),
	}
}

// Subscribe subscriber listen on topic
func (b *Broker) Subscribe(sub *Subscriber, topic string) bool {

	b.mu.Lock()
	defer b.mu.Unlock()

	if b.isEmpty(sub, topic) {
		return b.add(sub, topic)
	}

	if b.isExist(sub, topic) {
		return false
	}

	return b.add(sub, topic)
}

func (b *Broker) isExist(sub *Subscriber, topic string) bool {
	if len(b.subscribers[topic]) > 0 {
		for _, s := range b.subscribers[topic] {
			if sub == s {
				return true
			}
		}
	}
	return false
}

func (b *Broker) isEmpty(sub *Subscriber, topic string) bool {
	if len(b.subscribers[topic]) > 0 {
		return false
	}
	return true
}

// add adds new subscriber to listen on topic
func (b *Broker) add(sub *Subscriber, topic string) bool {
	b.subscribers[topic] = append(b.subscribers[topic], sub)
	return true
}

// Unsubscribe removes subscribers
func (b *Broker) Unsubscribe(sub *Subscriber, topic string) bool {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.remove(sub, topic)
}

// remove removes a subscriber from listen on topic
func (b *Broker) remove(sub *Subscriber, topic string) bool {
	for i, subscriber := range b.subscribers[topic] {
		if subscriber == sub {
			b.subscribers[topic] = append(b.subscribers[topic][:i], b.subscribers[topic][i+1:]...)
			return true
		}
	}
	return false
}

// Publish send message to channel
func (b *Broker) Publish(topic, msg string) {
	b.mu.Lock()
	for _, sub := range b.subscribers[topic] {
		sub.Do(msg)
	}
	b.mu.Unlock()
}

package pubsub

// Subscriber is subscriber
type Subscriber struct {
	ID         int
	handleFunc func(msg string)
}

// NewSubscriber new subscriber instance
func NewSubscriber() *Subscriber {
	return &Subscriber{}
}

func NewSubscriberWithFunc(handlerFunc func(msg string)) *Subscriber {
	return &Subscriber{handleFunc: handlerFunc}
}

func (s *Subscriber) SetHandleFunc(handlerFunc func(msg string)) {
	s.handleFunc = handlerFunc
}

func (s *Subscriber) Do(msg string) {
	s.handleFunc(msg)
}

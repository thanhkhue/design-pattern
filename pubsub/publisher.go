package pubsub

// Publisher stores record of subscriber and topics
type Publisher struct{}

// NewPublisher new publisher
func NewPublisher() *Publisher {
	return &Publisher{}
}

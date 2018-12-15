package publisher

import "time"

type PublisherI interface {
	Publish(topic string, message []byte) error
	MultiPublish(topic string, messages [][]byte) error
	DeferredPublish(topic string, delay time.Duration, message []byte) error
	Stop()
}

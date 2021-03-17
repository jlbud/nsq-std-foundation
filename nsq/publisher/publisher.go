package publisher

import "time"

// 生产者规则，生产者实例必须实现
type PublisherI interface {
	Publish(topic string, message []byte) error
	MultiPublish(topic string, messages [][]byte) error
	DeferredPublish(topic string, delay time.Duration, message []byte) error
	Stop()
}

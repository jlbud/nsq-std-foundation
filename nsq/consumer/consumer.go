package consumer

import "github.com/klbud/nsq-std-foundation/nsq/message"

// 消费者规则，消费者实例必须实现
type ConsumerI interface {
	SetParam(option string, value interface{})
	Start(handler message.HandlerI) error
	Stop()
}

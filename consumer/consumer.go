package consumer

import "github.com/klbud/nsq-std-foundation/message"

type ConsumerI interface {
	SetParam(option string, value interface{})
	Start(handler message.HandlerI) error
	Stop()
}

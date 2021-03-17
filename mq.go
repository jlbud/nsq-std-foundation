package mq

import (
	"time"
)

type MQType uint8

const (
	_ MQType = iota
	NSQ
)

type Pub interface {
	// Init means define method
	// without making connection or config check
	Init(topic string) error
	Publish(msg *Message) error
}

type Sub interface {
	// Init means define method
	// without making connection or config check
	Init(topic string, broadcastChannel string) error
	// handler include recover and execute in ONLY-ONE goroutine
	// handler has expired limit
	// error == nil means handle message success
	// error != nil means message will re-in queue
	Subscribe(handler func(*Message) error) error
	Stop() error
}

type Message struct {
	Content []byte
	// in pub, define delay publish
	// in sub, define delay re-in queue // handler's return != nil
	Delay time.Duration
	// the number of times this message was attempted
	Attempts uint16
}

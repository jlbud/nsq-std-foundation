package message

import (
	"time"

	"github.com/nsqio/go-nsq"
)

const MsgIDLength = 16

type MessageID [MsgIDLength]byte

// just for mock
func newMessage(id MessageID, body []byte) *Message {
	return nil
}

type MessageI interface {
	ID() [MsgIDLength]byte //消息ID
	Body() []byte          //待消费数据
	Attempts() uint16      //此消息被重复消费次数
	Finish()               //完成消费，默认自动完成

	Touch()                      //重制消息消费时间
	Requeue(delay time.Duration) //重新放回队列
}

type DependNsqMessage struct {
	msg *nsq.Message
}

type Message struct {
	DependNsqMessage
	Content   []byte
	TopicName string
}

func (m *Message) SetMsg(msg *nsq.Message) {
	m.msg = msg
}

func (m *Message) ID() [MsgIDLength]byte {
	return m.msg.ID
}

func (m *Message) Body() []byte {
	return m.msg.Body
}

func (m *Message) Attempts() uint16 {
	return m.msg.Attempts
}

func (m *Message) Finish() {
	m.msg.Finish()
}

func (m *Message) Requeue(delay time.Duration) {
	m.msg.Requeue(delay)
}

func (m *Message) Touch() {
	m.msg.Touch()
}

func (m *Message) RequeueWithoutBackoff(delay time.Duration) {
	m.msg.RequeueWithoutBackoff(delay)
}

type HandlerI interface {
	HandleMessage(message *Message) error
}

type HandlerFunc func(message *Message) error

func (h HandlerFunc) HandleMessage(m *Message) error {
	return h(m)
}
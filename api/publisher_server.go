package nsqkit

import (
	"fmt"

	"github.com/klbud/nsq-std-foundation/message"
	"github.com/klbud/nsq-std-foundation/publisher"
)

type PublisherServer struct {
	pi publisher.PublisherI
}

func NewPublisher() PublisherServerI {
	pinew, err := publisher.NewPublisher()
	if err != nil {
		panic(fmt.Errorf("NewPublisher error, %v", err))
	}
	ps := &PublisherServer{
		pi: pinew,
	}
	return ps
}

func (ps *PublisherServer) Publish(msg *message.Message) error {
	err := ps.pi.Publish(msg.TopicName, msg.Content)
	if err != nil {
		return err
	}
	return nil
}

func (ps *PublisherServer) StopPublisher() {
	ps.pi.Stop()
}

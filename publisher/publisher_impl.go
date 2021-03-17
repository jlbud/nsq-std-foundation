package publisher

import (
	"fmt"
	"strconv"
	"time"

	"github.com/klbud/nsq-std-foundation/mq"
	"github.com/nsqio/go-nsq"
)

type PublisherImpl struct {
	producer *nsq.Producer
}

func NewPublisher() (*PublisherImpl, error) {
	ndUrl := mq.GetConnNd().BroadcastAddress + ":" + strconv.Itoa(mq.GetConnNd().HostTcpPort)
	fmt.Println("NewPublisher", ndUrl)
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(ndUrl, config)
	if err != nil {
		return nil, err
	}
	return &PublisherImpl{producer: p}, nil
}

func (p *PublisherImpl) Publish(topic string, message []byte) error {
	err := p.producer.Publish(topic, message)
	if err != nil {
		return err
	}
	return nil
}

func (p *PublisherImpl) DeferredPublish(topic string, delay time.Duration, message []byte) error {
	if err := p.producer.DeferredPublish(topic, delay, message); err != nil {
		return err
	}
	return nil
}

func (p *PublisherImpl) MultiPublish(topic string, messages [][]byte) error {
	err := p.producer.MultiPublish(topic, messages)
	if err != nil {
		return err
	}
	return nil
}

func (p *PublisherImpl) Stop() {
	if p.producer != nil {
		p.producer.Stop()
	}
}

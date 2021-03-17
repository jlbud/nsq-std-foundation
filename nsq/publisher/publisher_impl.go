package publisher

import (
	"fmt"
	"strconv"
	"time"

	"github.com/klbud/nsq-std-foundation/nsq/common"
	"github.com/nsqio/go-nsq"
)

type PublisherImpl struct {
	producer *nsq.Producer
}

// 初始化生产者对象
func NewPublisher() (*PublisherImpl, error) {
	err := common.InitNsq()
	if err != nil {
		return nil, err
	}
	// nsqd tcp端口
	ndUrl := common.GetConnNd().BroadcastAddress + ":" + strconv.Itoa(common.GetConnNd().HostTcpPort)
	config := nsq.NewConfig()
	p, err := nsq.NewProducer(ndUrl, config)
	if err != nil {
		return nil, err
	}
	return &PublisherImpl{producer: p}, nil
}

// 推送消息
func (p *PublisherImpl) Publish(topic string, message []byte) error {
	if err := p.producer.Publish(topic, message); err != nil {
		// 发送消息错误后重连
		errConn := p.reConn()
		return fmt.Errorf("publish err %v, reconn err %v", err, errConn)
	}
	return nil
}

// 重新获取可用nsqd
func (p *PublisherImpl) reConn() error {
	err := common.InitNsq()
	if err != nil {
		return err
	}
	// nsqd tcp端口
	ndUrl := common.GetConnNd().BroadcastAddress + ":" + strconv.Itoa(common.GetConnNd().HostTcpPort)
	config := nsq.NewConfig()
	prod, err := nsq.NewProducer(ndUrl, config)
	if err != nil {
		return err
	}
	p.producer = prod
	return nil
}

// 延迟推送消息
func (p *PublisherImpl) DeferredPublish(topic string, delay time.Duration, message []byte) error {
	if err := p.producer.DeferredPublish(topic, delay, message); err != nil {
		errConn := p.reConn()
		return fmt.Errorf("deferredpublish err %v, reconn err %v", err, errConn)
	}
	return nil
}

// 推送多条消息
func (p *PublisherImpl) MultiPublish(topic string, messages [][]byte) error {
	err := p.producer.MultiPublish(topic, messages)
	if err != nil {
		return err
	}
	return nil
}

// 停止生产方服务
func (p *PublisherImpl) Stop() {
	if p.producer != nil {
		p.producer.Stop()
	}
}

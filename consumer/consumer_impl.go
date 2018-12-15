package consumer

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/Kevin005/nsq-std-foundation/config"
	"github.com/Kevin005/nsq-std-foundation/message"
	"github.com/nsqio/go-nsq"
)

type logger interface {
	Output(int, string) error
}

type ConsumerImpl struct {
	consumer    *nsq.Consumer
	config      *nsq.Config
	concurrency int
	channel     string
	topic       string
	level       nsq.LogLevel
	log         logger
	err         error
}

func NewConsumer() (error, ConsumerI) {
	config := nsq.NewConfig()
	config.MaxBackoffDuration = 20 * time.Millisecond        //处理失败时退回的最长时间
	config.LookupdPollInterval = 1000 * time.Millisecond     //设置重连时间
	config.RDYRedistributeInterval = 1000 * time.Millisecond //将max-in-flight重新分配到连接之间的持续时间
	config.AuthSecret = ""                                   //认证密钥,暂时不需要
	return nil, &ConsumerImpl{
		config: config,
		log:    log.New(os.Stderr, "", log.LstdFlags),
		level:  nsq.LogLevelInfo,
	}
}

func (c *ConsumerImpl) Start(h message.HandlerI) error {
	if c.err != nil {
		return c.err
	}

	client, err := nsq.NewConsumer(c.topic, c.channel, c.config)
	if err != nil {
		return err
	}
	c.consumer = client
	client.SetLogger(c.log, c.level)
	client.AddConcurrentHandlers(nsq.HandlerFunc(func(msg *nsq.Message) error {
		m := &message.Message{}
		m.SetMsg(msg)
		err := h.HandleMessage(m)
		return err
	}), c.concurrency)
	return c.connect()
}

func (c *ConsumerImpl) connect() error {
	if len(config.LookupdsAddress) > 0 {
		err := c.consumer.ConnectToNSQLookupds(config.LookupdsAddress)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("lookupds address empty")
}

func (c *ConsumerImpl) SetLogger(log logger, level nsq.LogLevel) {
	c.level = level
	c.log = log
}

func (c *ConsumerImpl) setMap(options map[string]interface{}) {
	for k, v := range options {
		c.SetParam(k, v)
	}
}

func (c *ConsumerImpl) SetParam(option string, value interface{}) {
	switch option {
	case "channel":
		c.channel = value.(string)
	case "topic":
		c.topic = value.(string)
	case "concurrency":
		c.concurrency = value.(int)
	default:
		err := c.config.Set(option, value)
		if err != nil {
			c.err = err
		}
	}
}

func (c *ConsumerImpl) Stop() {
	c.consumer.Stop()
	<-c.consumer.StopChan
}

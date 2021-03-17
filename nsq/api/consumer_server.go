package nsqkit

import (
	"fmt"
	"time"

	"github.com/klbud/nsq-std-foundation/nsq/consumer"
	"github.com/klbud/nsq-std-foundation/nsq/message"
	"github.com/pkg/errors"
)

type ConsumerServer struct {
	ci consumer.ConsumerI
}

// todo 可以用反射实现tag解析
type ConsumerConfigP struct {
	MaxAttempts         int           `opt:"max_attempts" default:"10"`                     //消息最大重发次数，超过则废弃
	MaxInFlight         int           `opt:"max_in_flight" default:"2500"`                  //此连接上传输中的最大消息数，飞行中允许的最大消息数
	MsgTimeout          time.Duration `opt:"msg_timeout" default:"5*time.Second"`           //传递给此客户端的消息的服务器端消息超时
	DefaultRequeueDelay time.Duration `opt:"default_requeue_delay" default:"1*time.Second"` //消息重发延迟时间
}

type ConsumerP struct {
	*ConsumerConfigP
	TopicName   string
	ChannelName string
}

func NewConsumer(cp *ConsumerP) ConsumerServerI {
	err, c := consumer.NewConsumer()
	if err != nil {
		panic(fmt.Errorf("NewConsumer error, %v", err))
	}
	err, cp = cp.checkConsumerP(cp)
	if err != nil {
		panic(fmt.Errorf("checkConsumerP error, %v", err))
	}
	c.SetParam("channel", cp.ChannelName)
	c.SetParam("topic", cp.TopicName)
	c.SetParam("max_attempts", cp.MaxAttempts)
	c.SetParam("max_in_flight", cp.MaxInFlight)
	c.SetParam("default_requeue_delay", cp.DefaultRequeueDelay)
	c.SetParam("msg_timeout", cp.MsgTimeout)
	cs := &ConsumerServer{
		ci: c,
	}
	return cs
}

// todo 可以用用反射实现tag解析
func (_ *ConsumerP) checkConsumerP(cp *ConsumerP) (error, *ConsumerP) {
	if cp.TopicName == "" || cp.ChannelName == "" {
		return errors.New("neither topicname and channelname can be empty"), nil
	}
	if cp.ConsumerConfigP == nil {
		cp.ConsumerConfigP = &ConsumerConfigP{
			DefaultRequeueDelay: 1 * time.Second,
			MsgTimeout:          5 * time.Second,
			MaxAttempts:         10,
			MaxInFlight:         2500,
		}
		return nil, cp
	}
	if cp.DefaultRequeueDelay == 0 {
		cp.DefaultRequeueDelay = 1 * time.Second
	}
	if cp.MsgTimeout == 0 {
		cp.MsgTimeout = 5 * time.Second
	}
	if cp.MaxAttempts == 0 {
		cp.MaxAttempts = 10
	}
	if cp.MaxInFlight == 0 {
		cp.MaxInFlight = 1
	}
	return nil, cp
}

func (cs *ConsumerServer) StartHandler(handler message.HandlerI) error {
	cs.ci.SetParam("concurrency", 1)
	if err := cs.ci.Start(handler); err != nil {
		return err
	}
	return nil
}

func (cs *ConsumerServer) StartConcurrentHandlers(handler message.HandlerI) error {
	cs.ci.SetParam("concurrency", 15)
	if err := cs.ci.Start(handler); err != nil {
		return err
	}
	return nil
}

func (cs *ConsumerServer) StopConsumer() {
	cs.ci.Stop()
}

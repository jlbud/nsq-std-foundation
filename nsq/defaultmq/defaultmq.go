package defaultmq

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/klbud/nsq-std-foundation"
	"github.com/klbud/nsq-std-foundation/nsq/common"
	nsqConfig "github.com/klbud/nsq-std-foundation/nsq/config"
	"github.com/klbud/nsq-std-foundation/nsq/consumer"
	"github.com/klbud/nsq-std-foundation/nsq/message"
	"github.com/klbud/nsq-std-foundation/nsq/publisher"
)

// 实现了/mq/mq.go标准
// 根据此标准提供服务

/////////////////////////////////////// Pub start
// 生成publisher实例
func NewPub(errorLog func(log string) error, config map[string]interface{}) (pub mq.Pub, err error) {
	defer func() {
		if errNewPub := recover(); errNewPub != nil {
			err = fmt.Errorf("newpub err %v", errNewPub)
		}
	}()
	pp := &publisherP{}
	if err := pp.initParameter(config); err != nil {
		return nil, err
	}
	p, err := publisher.NewPublisher()
	if err != nil {
		return nil, err
	}
	pn := &PubNSQ{}
	pn.pi = p
	pn.errLog = errorLog
	return pn, nil
}

// publisher模版
type PubNSQ struct {
	errLog    func(log string) error
	pi        publisher.PublisherI
	topicName string
}

// 初始化publisher实例topic
func (p *PubNSQ) Init(topic string) error {
	if topic == "" {
		return errors.New("[topic] not be empty")
	}
	p.topicName = topic
	return nil
}

// publisher实例生产一条消息
func (p *PubNSQ) Publish(msg *mq.Message) (err error) {
	defer func() {
		if errPublish := recover(); errPublish != nil {
			err = fmt.Errorf("publish err %v", errPublish)
		}
	}()
	if msg.Delay != 0 {
		if err := p.pi.DeferredPublish(p.topicName, msg.Delay, msg.Content); err != nil {
			p.errLog(err.Error())
			return err
		}
	} else {
		if err := p.pi.Publish(p.topicName, msg.Content); err != nil {
			p.errLog(err.Error())
			return err
		}
	}
	return nil
}

type publisherP struct{}

// 初始化publisher参数
func (_ *publisherP) initParameter(config map[string]interface{}) (err error) {
	defer func() {
		if initPubErr := recover(); initPubErr != nil {
			err = fmt.Errorf("newpub initparameter err %v", initPubErr)
		}
	}()
	if v, ok := config["lookupds_address"]; ok {
		nsqConfig.LookupdsAddress = v.([]string)
	} else {
		nsqConfig.LookupdsAddress = []string{common.LOOKUPD_ADDRESS_153, common.LOOKUPD_ADDRESS_154}
	}
	return nil
}

/////////////////////////////////////// Pub end

/////////////////////////////////////// Sub start
// 生成consumer实例
func NewSub(errorLog func(log string) error, config map[string]interface{}) (sub mq.Sub, err error) {
	defer func() {
		if errNewSub := recover(); errNewSub != nil {
			err = fmt.Errorf("newsub err %v", errNewSub)
		}
	}()
	err, c := consumer.NewConsumer()
	if err != nil {
		return nil, err
	}
	cp := &consumerP{}
	if err := cp.initParameter(config); err != nil {
		return nil, err
	}
	sn := &SubNSQ{}
	sn.errLog = errorLog
	sn.cp = cp
	sn.ci = c
	return sn, nil
}

// consumer模版
type SubNSQ struct {
	errLog func(log string) error
	// 消费者
	ci consumer.ConsumerI
	// 消费者配置
	cp *consumerP
}

// 初始化consumer实例的参数
func (s *SubNSQ) Init(topic string, broadcastChannel string) error {
	if topic == "" || broadcastChannel == "" {
		return errors.New("[topic] and [broadcastChannel] not be empty")
	}
	s.ci.SetParam("topic", topic)
	s.ci.SetParam("channel", broadcastChannel)
	s.ci.SetParam("concurrency", s.cp.concurrency)
	s.ci.SetParam("max_attempts", s.cp.maxAttempts)
	s.ci.SetParam("max_in_flight", s.cp.maxInFlight)
	s.ci.SetParam("default_requeue_delay", s.cp.defaultRequeueDelay)
	s.ci.SetParam("msg_timeout", s.cp.msgTimeout)
	return nil
}

// consumer实例订阅消息
// handler为注册对象使用的回调函数
func (s *SubNSQ) Subscribe(handler func(*mq.Message) error) error {
	startErr := s.ci.Start(message.HandlerFunc(func(message *message.Message) error {
		defer func() {
			if err := recover(); err != nil {
				buf := make([]byte, 1<<20)
				num := runtime.Stack(buf, false)
				s.errLog(fmt.Sprintf("%v %v %v", err, num, string(buf)))
				s.errLog(fmt.Errorf("subscribe processing message err: %v", err).Error())
			}
		}()
		if message == nil || message.Body() == nil {
			s.errLog(fmt.Errorf("subscribe processing message err: message is empty").Error())
			return nil
		}
		m := &mq.Message{
			Content:  message.Body(),
			Attempts: message.Attempts(),
		}
		err := handler(m)
		if err != nil && m.Delay != 0 {
			message.Requeue(m.Delay)
		}
		return err
	}))
	return startErr
}

// consumer停止监听消息
func (s *SubNSQ) Stop() error {
	s.ci.Stop()
	return nil
}

//todo 需要反射
type consumerP struct {
	concurrency         int           `opt:"concurrency" default:"1"`                       // 处理消息的回调线程为1
	maxAttempts         int           `opt:"max_attempts" default:"10"`                     // 消息最大重发次数，超过则废弃
	maxInFlight         int           `opt:"max_in_flight" default:"2500"`                  // 此连接上传输中的最大消息数，飞行中允许的最大消息数
	msgTimeout          time.Duration `opt:"msg_timeout" default:"5*time.Second"`           // 传递给此客户端的消息的服务器端消息超时
	defaultRequeueDelay time.Duration `opt:"default_requeue_delay" default:"1*time.Second"` // 消息重发延迟时间
}

func (cp *consumerP) initParameter(config map[string]interface{}) (err error) {
	defer func() {
		if initSubErr := recover(); initSubErr != nil {
			err = fmt.Errorf("newsub initparameter err %v", initSubErr)
		}
	}()
	if config == nil || len(config) == 0 {
		return errors.New("newsub initparameter err, [config] not be empty")
	}

	if v, ok := config["lookupds_address"]; ok {
		nsqConfig.LookupdsAddress = v.([]string)
	} else {
		nsqConfig.LookupdsAddress = []string{common.LOOKUPD_ADDRESS_153, common.LOOKUPD_ADDRESS_154}
	}

	if v, ok := config["max_attempts"]; ok {
		cp.maxAttempts = v.(int)
	} else {
		cp.maxAttempts = 10
	}

	if v, ok := config["max_in_flight"]; ok {
		cp.maxInFlight = v.(int)
	} else {
		cp.maxInFlight = 2500
	}

	if v, ok := config["msg_timeout"]; ok {
		cp.msgTimeout = v.(time.Duration)
	} else {
		cp.msgTimeout = 5 * time.Second
	}

	if v, ok := config["default_requeue_delay"]; ok {
		cp.defaultRequeueDelay = v.(time.Duration)
	} else {
		cp.defaultRequeueDelay = 60 * time.Second
	}

	if v, ok := config["concurrency"]; ok {
		cp.concurrency = v.(int)
	} else {
		cp.concurrency = 1
	}
	return nil
}

/////////////////////////////////////// Sub end

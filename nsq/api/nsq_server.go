package nsqkit

import (
	"github.com/klbud/nsq-std-foundation/nsq/message"
)

type PublisherServerI interface {
	Publish(msg *message.Message) error
	StopPublisher()
}

type ConsumerServerI interface {
	StartHandler(handler message.HandlerI) error
	StartConcurrentHandlers(handler message.HandlerI) error
	StopConsumer()
}

type CommonServerI interface {
	CreateTopic(topicName string) (error, *TopicResponse)
	//DeleteTopic(topicName string) (error, *TopicResponse) //删除topic，对外敏感
	GetAllTopics() (error, *TopicsNameResponse)
	CreateChannel(topicName, channelName string) (error, *ChannelResponse)
	//DeleteChannel(topicName, channelName string) (error, *ChannelResponse) //删除channel，对外敏感
	GetChannelsOfTopic(topicName string) (error, *ChannelsNameResponse)
}

package nsqkit

import (
	"encoding/json"
	"log"

	"github.com/Kevin005/nsq-std-foundation/mq/channels"
	"github.com/Kevin005/nsq-std-foundation/mq/topics"
)

func NewCommonServer() CommonServerI {
	return &CommonServer{
		topics:   topics.NewTopic(),
		channels: channels.NewChannel(),
	}
}

type CommonServer struct {
	topics   topics.TopicI
	channels channels.ChannelI
}

type TopicResponse struct {
	TopicName string
}

type ChannelResponse struct {
	TopicName   string
	ChannelName string
}

type TopicsNameResponse struct {
	Topics []string `json:topics`
}

type ChannelsNameResponse struct {
	Channels []string `json:channels`
}

func (cos *CommonServer) CreateTopic(topicName string) (error, *TopicResponse) {
	log.Printf("nsq create topic %v", topicName)
	err := cos.topics.CreateTopic(topicName)
	if err != nil {
		return err, nil
	}
	tr := &TopicResponse{
		TopicName: topicName,
	}
	return nil, tr
}

func (cos *CommonServer) DeleteTopic(topicName string) (error, *TopicResponse) {
	log.Printf("nsq delete topic %v", topicName)
	err := cos.topics.DeleteTopic(topicName)
	if err != nil {
		return err, nil
	}
	tr := &TopicResponse{
		TopicName: topicName,
	}
	return nil, tr
}

func (cos *CommonServer) GetAllTopics() (error, *TopicsNameResponse) {
	log.Printf("nsq get all topic")
	err, res := cos.topics.GetTopics()
	if err != nil {
		return err, nil
	}
	tr := &TopicsNameResponse{}
	err = json.Unmarshal([]byte(res), tr)
	if err != nil {
		return err, nil
	}
	return nil, tr
}

func (cos *CommonServer) CreateChannel(topicName, channelName string) (error, *ChannelResponse) {
	log.Printf("nsq create channel, topic name is %v, channel name is %v ", topicName, channelName)
	err := cos.channels.CreateChannel(topicName, channelName)
	if err != nil {
		return err, nil
	}
	cr := &ChannelResponse{
		TopicName:   topicName,
		ChannelName: channelName,
	}
	return nil, cr
}

func (cos *CommonServer) DeleteChannel(topicName, channelName string) (error, *ChannelResponse) {
	log.Printf("nsq delete channel %v %v", topicName, channelName)
	err := cos.channels.DeleteChannel(topicName, channelName)
	if err != nil {
		return err, nil
	}
	cr := &ChannelResponse{
		TopicName:   topicName,
		ChannelName: channelName,
	}
	return nil, cr
}

func (cos *CommonServer) GetChannelsOfTopic(topicName string) (error, *ChannelsNameResponse) {
	log.Printf("nsq get channels of topic %v", topicName)
	err, res := cos.channels.GetChannelsOfTopic(topicName)
	if err != nil {
		return err, nil
	}
	cr := &ChannelsNameResponse{}
	err = json.Unmarshal([]byte(res), cr)
	if err != nil {
		return err, nil
	}
	return nil, cr
}

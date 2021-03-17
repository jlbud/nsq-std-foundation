package nsqkit

import (
	"encoding/json"

	"github.com/klbud/nsq-std-foundation/nsq/common/channels"
	"github.com/klbud/nsq-std-foundation/nsq/common/topics"
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

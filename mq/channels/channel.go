package channels

type ChannelI interface {
	CreateChannel(topicName, channelName string) error
	DeleteChannel(topicName, channelName string) error
	EmptyChannel(topicName, channelName string) error
	PauseChannel(topicName, channelName string) error
	UnpauseChannel(topicName, channelName string) error
	GetChannelsOfTopic(topicName string) (error,string)
}

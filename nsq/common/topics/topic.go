package topics

type TopicI interface {
	CreateTopic(topicName string) error
	DeleteTopic(topicName string) error
	EmptyTopic(topicName string) error
	PauseTopic(topicName string) error
	UnpauseTopic(topicName string) error
	GetTopics() (error, string)
}

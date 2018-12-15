
##
提供了一个最简化的NSQ封装

### 目录结构
```
+-- api
|   +-- nsq_server        //对外总线
|   +-- consumer_server   //对外消费方服务
|   +-- publisher_server  //对外生产方服务
|   +-- common_server     //对外公共服务(增加、删除topic channel等等)
|
+-- publisher             //生产方服务
|   +-- publisher         //生产方接口
|   +-- publisher_impl    //生产方接口实现
|
+-- consumer              //消费方服务
|   +-- consumer          //消费方接口
|   +-- consumer_impl     //消费方接口实现
|
+-- mq                    //消息队列服务
|   +-- topics            //topic相关
|   +-- channels          //channel相关
|   +-- lookupd           //服务发现相关
|
+-- message               //对外消息数据结构
|
+-- utils                 //工具包
```

### 对外接口服务
``` golang
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
	DeleteTopic(topicName string) (error, *TopicResponse) //删除topic，对外敏感
	GetAllTopics() (error, *TopicsNameResponse)
	CreateChannel(topicName, channelName string) (error, *ChannelResponse)
	DeleteChannel(topicName, channelName string) (error, *ChannelResponse) //删除channel，对外敏感
	GetChannelsOfTopic(topicName string) (error, *ChannelsNameResponse)
}
```

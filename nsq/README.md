##
提供了一个NSQ封装

### 目录结构
```
+-- api                   // 对外服务网关
|   +-- nsq_server        // 对外服务总线
|   +-- consumer_server   // 对外消费方服务
|   +-- publisher_server  // 对外生产方服务
|   +-- common_server     // 对外公共服务(增加、删除topic channel等等)
|
+-- publisher             // 生产方服务
|   +-- publisher         // 生产方接口
|   +-- publisher_impl    // 生产方接口实现
|
+-- consumer              // 消费方服务
|   +-- consumer          // 消费方接口
|   +-- consumer_impl     // 消费方接口实现
|
+-- common                // 公共服务
|   +-- topics            // topic创建、删除等
|   +-- channels          // channel创建、删除等
|   +-- lookupd           // 服务发现相关，获取可用nsqd
|   +-- mq_common         // 初始化可用nsqd服务，队列http接口列表
|
+-- defaultmq             // 实现了 /mq/mq.go标准，对外提供服务
|
+-- message               // 对外消息数据结构
|
+-- utils                 // 工具包

```
### 标准实现
* 如果nsq为独立项目，则使用nsq_server为标准初始化实例即可(common_server.go、publisher_server.go、consumer_server.go已实现)
* 如果nsq为mq中一员，则使用 /mq/mq.go为标准初始化实例即可(defaultmq.go已实现)

### api对外接口服务
``` golang
// 消费方服务
type PublisherServerI interface {
	Publish(msg *message.Message) error // 生产一条消息
	StopPublisher() // 停止生产
}

// 消费方服务
type ConsumerServerI interface {
	StartHandler(handler message.HandlerI) error // 开始监听消息
	StartConcurrentHandlers(handler message.HandlerI) error // 开始异步监听消息
	StopConsumer() // 停止监听
}

// 通用服务
type CommonServerI interface {
	CreateTopic(topicName string) (error, *TopicResponse) // 创建topic
	DeleteTopic(topicName string) (error, *TopicResponse) // 删除topic，对外敏感
	GetAllTopics() (error, *TopicsNameResponse) // 获取所有topic
	CreateChannel(topicName, channelName string) (error, *ChannelResponse) // 创建channel
	DeleteChannel(topicName, channelName string) (error, *ChannelResponse) // 删除channel，对外敏感
	GetChannelsOfTopic(topicName string) (error, *ChannelsNameResponse) // 获取某个topic下所有的channel
}
```

### 服务端口
| 服务类型 | 端口类型 | 端口号 |
| ------ | ------ | ------ |
| lookupd | http | 4161 |
| lookupd | tcp | 4160 |
| nsqd | http | 4151 |
| nsqd | tcp | 4150 |
| nsqadmin | http | 4171 |

### 项目文档
* 调研文档
* 部署文档
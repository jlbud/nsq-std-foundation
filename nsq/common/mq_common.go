package common

import (
	"errors"
	"strconv"

	"github.com/klbud/nsq-std-foundation/nsq/config"
)

var (
	topic_create  = "/topic/create"  //创建topic
	topic_delete  = "/topic/delete"  //删除topic
	topic_empty   = "/topic/empty"   //清空topic队列里的信息
	topic_pause   = "/topic/pause"   //暂停消息往topic的所有channel，消息将按主题排队
	topic_unpause = "/topic/unpause" //恢复topic里暂停的消息
	topic_lists   = "/topics"        //查看所有的topics

	channel_create  = "/channel/create"  //创建channel
	channel_delete  = "/channel/delete"  //删除channel
	channel_empty   = "/channel/empty"   //清空channel里所有排队的信息
	channel_pause   = "/channel/pause"   //暂停消费者消费channel里的消息
	channel_unpause = "/channel/unpause" //恢复消息
	channel_lists   = "/channels"        //查看一个topic的所有channel

	all_nsqds = "/nodes" // 获取lookupd中所有的nsqd地址
)

var (
	nd *CurrentNsqd // nsqd地址
)

var (
	// 生产服务器nsq服务发现地址
	LOOKUPD_ADDRESS_153 = "192.168.7.153:4161"
	LOOKUPD_ADDRESS_154 = "192.168.7.154:4161"
)

// 当前可以连接的nsqd
type CurrentNsqd struct {
	HostHttpPort     int
	HostTcpPort      int
	HostName         string // 格式：xiaomingdeMacBook-Air.local
	RemoteAddress    string // 格式：127.0.0.1:51023
	BroadcastAddress string // nsqd ip地址
}

func CreateTopicUrl(topicName string) string {
	return "http://" + nd.BroadcastAddress + ":" + strconv.Itoa(nd.HostHttpPort) + topic_create + "?topic=" + topicName
}

func DeleteTopicUrl(topicName string) string {
	return "http://" + nd.BroadcastAddress + ":" + strconv.Itoa(nd.HostHttpPort) + topic_delete + "?topic=" + topicName
}

func EmptyTopicUrl(topicName string) string {
	return "http://" + nd.BroadcastAddress + ":" + strconv.Itoa(nd.HostHttpPort) + topic_empty + "?topic=" + topicName
}

func PauseTopicUrl(topicName string) string {
	return "http://" + nd.BroadcastAddress + ":" + strconv.Itoa(nd.HostHttpPort) + topic_pause + "?topic=" + topicName
}

func UnpauseTopicUrl(topicName string) string {
	return "http://" + nd.BroadcastAddress + ":" + strconv.Itoa(nd.HostHttpPort) + topic_unpause + "?topic=" + topicName
}

func CreateChannelUrl(topicName, channelName string) string {
	return "http://" + nd.BroadcastAddress + ":" + strconv.Itoa(nd.HostHttpPort) + channel_create + "?topic=" + topicName + "&channel=" + channelName
}

func DeleteChannelUrl(topicName, channelName string) string {
	return "http://" + nd.BroadcastAddress + ":" + strconv.Itoa(nd.HostHttpPort) + channel_delete + "?topic=" + topicName + "&channel=" + channelName
}

func EmptyChannelUrl(topicName, channelName string) string {
	return "http://" + nd.BroadcastAddress + ":" + strconv.Itoa(nd.HostHttpPort) + channel_empty + "?topic=" + topicName + "&channel=" + channelName
}

func PauseChannelUrl(topicName, channelName string) string {
	return "http://" + nd.BroadcastAddress + ":" + strconv.Itoa(nd.HostHttpPort) + channel_pause + "?topic=" + topicName + "&channel=" + channelName
}

func UnpauseChannelUrl(topicName, channelName string) string {
	return "http://" + nd.BroadcastAddress + ":" + strconv.Itoa(nd.HostHttpPort) + channel_unpause + "?topic=" + topicName + "&channel=" + channelName
}

func GetAllNsqdsUrl() string {
	return all_nsqds
}

func GetAllTopics() string {
	return topic_lists
}

func GetChannelsOfTopic() string {
	return channel_lists
}

// 初始化可以连接的nsqd地址
func initConnNsqdAddress() error {
	var httpLookups []string
	for _, l := range config.LookupdsAddress {
		httpLookups = append(httpLookups, "http://"+l)
	}
	ld := &Lookupd{
		LookupdAddress: httpLookups,
	}
	err, nds := ld.GetNsqds()
	if err != nil {
		return err
	}

	for _, pd := range nds.Producers {
		nd = &CurrentNsqd{
			HostHttpPort:     pd.HttpPort,
			HostTcpPort:      pd.TcpPort,
			RemoteAddress:    pd.RemoteAddress,
			BroadcastAddress: pd.BroadcastAddress,
		}
		break
	}
	if nd == nil {
		return errors.New("initConnNsqdAddress fail, nsqd is not found by lookup address")
	}
	return nil
}

func InitNsq() error {
	err := initConnNsqdAddress()
	if err != nil {
		return err
	}
	return nil
}

// 当前可用的nsqd
func GetConnNd() *CurrentNsqd {
	return nd
}

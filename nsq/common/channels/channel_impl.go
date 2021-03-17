package channels

import (
	"errors"

	"github.com/klbud/nsq-std-foundation/nsq/common"
	"github.com/klbud/nsq-std-foundation/nsq/config"
	"github.com/klbud/nsq-std-foundation/nsq/util/http"
)

type ChannelImpl struct{}

func NewChannel() ChannelI {
	return &ChannelImpl{}
}

func (_ *ChannelImpl) CreateChannel(topicName, channelName string) error {
	hr := &http.HttpResult{}
	er := http.SendPost(nil, common.CreateChannelUrl(topicName, channelName), hr)
	if er != nil {
		return er
	}
	if hr.StateCode == http.HttpResultSuccess {
		return nil
	} else {
		return errors.New(hr.Body)
	}
}

func (_ *ChannelImpl) DeleteChannel(topicName, channelName string) error {
	hr := &http.HttpResult{}
	er := http.SendPost(nil, common.DeleteChannelUrl(topicName, channelName), hr)
	if er != nil {
		return er
	}
	if hr.StateCode == http.HttpResultSuccess {
		return nil
	} else {
		return errors.New(hr.Body)
	}
}

func (_ *ChannelImpl) EmptyChannel(topicName, channelName string) error {
	hr := &http.HttpResult{}
	er := http.SendPost(nil, common.EmptyChannelUrl(topicName, channelName), hr)
	if er != nil {
		return er
	}
	if hr.StateCode == http.HttpResultSuccess {
		return nil
	} else {
		return errors.New(hr.Body)
	}
}

func (_ *ChannelImpl) PauseChannel(topicName, channelName string) error {
	hr := &http.HttpResult{}
	er := http.SendPost(nil, common.PauseChannelUrl(topicName, channelName), hr)
	if er != nil {
		return er
	}
	if hr.StateCode == http.HttpResultSuccess {
		return nil
	} else {
		return errors.New(hr.Body)
	}
}

func (_ *ChannelImpl) UnpauseChannel(topicName, channelName string) error {
	hr := &http.HttpResult{}
	er := http.SendPost(nil, common.UnpauseChannelUrl(topicName, channelName), hr)
	if er != nil {
		return er
	}
	if hr.StateCode == http.HttpResultSuccess {
		return nil
	} else {
		return errors.New(hr.Body)
	}
}

func (_ *ChannelImpl) GetChannelsOfTopic(topicName string) (error, string) {
	hr := &http.HttpResult{}
	var httpLookups []string
	for _, l := range config.LookupdsAddress {
		httpLookups = append(httpLookups, "http://"+l)
	}
	if httpLookups == nil || len(httpLookups) == 0 {
		return errors.New("lookupdAddress empty"), ""
	}

	for _, lkaddress := range httpLookups {
		er := http.SendGet(lkaddress+common.GetChannelsOfTopic()+"?topic="+topicName, hr)
		if er != nil {
			return er, ""
		}
		if hr.StateCode == http.HttpResultSuccess {
			return nil, hr.Body
		} else {
			return errors.New(hr.Body), ""
		}
	}
	return nil, ""
}

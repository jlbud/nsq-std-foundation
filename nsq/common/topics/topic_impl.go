package topics

import (
	"errors"

	"github.com/klbud/nsq-std-foundation/nsq/common"
	"github.com/klbud/nsq-std-foundation/nsq/config"
	"github.com/klbud/nsq-std-foundation/nsq/util/http"
)

type TopicImpl struct{}

func NewTopic() TopicI {
	return &TopicImpl{}
}

func (t *TopicImpl) CreateTopic(topicName string) error {
	hr := &http.HttpResult{}
	er := http.SendPost(nil, common.CreateTopicUrl(topicName), hr)
	if er != nil {
		return er
	}
	if hr.StateCode == http.HttpResultSuccess {
		return nil
	} else {
		return errors.New(hr.Body)
	}
}

func (t *TopicImpl) DeleteTopic(topicName string) error {
	hr := &http.HttpResult{}
	err := http.SendPost(nil, common.DeleteTopicUrl(topicName), hr)
	if err != nil {
		return err
	}
	if hr.StateCode == http.HttpResultSuccess {
		return nil
	} else {
		return errors.New(hr.Body)
	}
}

func (t *TopicImpl) EmptyTopic(topicName string) error {
	hr := &http.HttpResult{}
	err := http.SendPost(nil, common.EmptyTopicUrl(topicName), hr)
	if err != nil {
		return err
	}
	if hr.StateCode == http.HttpResultSuccess {
		return nil
	} else {
		return errors.New(hr.Body)
	}
}

func (t *TopicImpl) PauseTopic(topicName string) error {
	hr := &http.HttpResult{}
	err := http.SendPost(nil, common.PauseTopicUrl(topicName), hr)
	if err != nil {
		return err
	}
	if hr.StateCode == http.HttpResultSuccess {
		return nil
	} else {
		return errors.New(hr.Body)
	}
}

func (t *TopicImpl) UnpauseTopic(topicName string) error {
	hr := &http.HttpResult{}
	err := http.SendPost(nil, common.UnpauseTopicUrl(topicName), hr)
	if err != nil {
		return err
	}
	if hr.StateCode == http.HttpResultSuccess {
		return nil
	} else {
		return errors.New(hr.Body)
	}
}

func (t *TopicImpl) GetTopics() (error, string) {
	hr := &http.HttpResult{}
	var httpLookups []string
	for _, l := range config.LookupdsAddress {
		httpLookups = append(httpLookups, "http://"+l)
	}
	if httpLookups == nil || len(httpLookups) == 0 {
		return errors.New("lookupdAddress empty"), ""
	}

	for _, lkaddress := range httpLookups {
		er := http.SendGet(lkaddress+common.GetAllTopics(), hr)
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

package nsqkit

import (
	"testing"
)

func TestCommonServer_CreateTopic(t *testing.T) {
	ncs := NewCommonServer()
	err, tr := ncs.CreateTopic("test3")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(tr.TopicName)
}

//func TestCommonServer_DeleteTopic(t *testing.T) {
//	ncs := NewCommonServer()
//	err, tr := ncs.DeleteTopic("test1")
//	if err != nil {
//		t.Error(err.Error())
//		return
//	}
//	t.Log(tr.TopicName)
//}

func TestCommonServer_CreateChannel(t *testing.T) {
	ncs := NewCommonServer()
	err, cr := ncs.CreateChannel("test3", "b")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(cr.TopicName, "\n", cr.ChannelName)
}

//func TestCommonServer_DeleteChannel(t *testing.T) {
//	ncs := NewCommonServer()
//	err, cr := ncs.DeleteChannel("test1", "b")
//	if err != nil {
//		t.Error(err.Error())
//		return
//	}
//	t.Log(cr.TopicName, "\n", cr.ChannelName)
//}

func TestCommonServer_GetAllTopics(t *testing.T) {
	ncs := NewCommonServer()
	err, cr := ncs.GetAllTopics()
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(cr.Topics)
}

func TestCommonServer_GetChannelsOfTopic(t *testing.T) {
	ncs := NewCommonServer()
	err, cr := ncs.GetChannelsOfTopic("test1")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(cr.Channels)
}

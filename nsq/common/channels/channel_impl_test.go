package channels

import (
	"testing"
)

func TestChannelImpl_CreateChannel(t *testing.T) {
	ci := NewChannel()
	err := ci.CreateChannel("test_topic6", "4")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("create channel success")
}

func TestChannelImpl_DeleteChannel(t *testing.T) {
	//get http://127.0.0.1:4151/channel/delete?topic=test&channel=d
	ci := NewChannel()
	err := ci.DeleteChannel("test", "1")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("delete channel success")
}

func TestChannelImpl_GetChannelsOfTopic(t *testing.T) {
	ci := NewChannel()
	err, res := ci.GetChannelsOfTopic("test")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("get success", res)
}

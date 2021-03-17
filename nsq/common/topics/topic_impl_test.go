package topics

import "testing"

func TestTopicImpl_CreateTopic(t *testing.T) {
	topicName := "test_topic5"
	ti := NewTopic()
	err := ti.CreateTopic(topicName)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("create success", topicName)
}

func TestTopicImpl_DeleteTopic(t *testing.T) {
	topicName := "test_topic4"
	ti := NewTopic()
	err := ti.DeleteTopic(topicName)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("delete success", topicName)
}

func TestTopicImpl_GetTopics(t *testing.T) {
	ti := NewTopic()
	err, res := ti.GetTopics()
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("get success", res)
}

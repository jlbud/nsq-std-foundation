package http

import (
	"testing"
)

func TestSendGet(t *testing.T) {
	r := &HttpResult{}
	err := SendGet("http://www.kuaidi100.com/query?type=yuantong&postid=11111111111", r)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(r.StateCode, r.Body)
}

func TestSendPost(t *testing.T) {
	r := &HttpResult{}
	err := SendPost(nil, "http://www.kuaidi100.com/query?type=yuantong&postid=11111111111", r)
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log(r.StateCode, r.Body)
}

func TestSendPostForm(t *testing.T) {}

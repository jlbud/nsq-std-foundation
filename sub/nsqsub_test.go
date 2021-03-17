package sub

import (
	"sync"
	"testing"
	"time"

	"github.com/klbud/nsq-std-foundation"
)

func TestSubNSQ(t *testing.T) {
	cfg := map[string]interface{}{
		"max_attempts":          10,
		"concurrency":           10,
		"default_requeue_delay": 1 * time.Second,
		"msg_timeout":           5 * time.Second,
		"lookupds_address":      []string{"192.168.1.174:4161"},
	}
	// 初始化nsq队列
	sub, err := New(mq.NSQ, func(log string) error {
		t.Errorf("err log %s", log)
		return nil
	}, cfg)
	if err != nil {
		t.Errorf("new sub err %v", err)
		return
	}

	// 消费端消费注册
	err = sub.Init("test", "default")
	if err != nil {
		t.Error(err)
		return
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		time.Sleep(5 * time.Second)
		// 消费端停止指令
		sub.Stop()
		wg.Done()
	}()

	count := 0
	beforeTime := time.Now()
	// 消费端开始消费
	err = sub.Subscribe(func(msg *mq.Message) error {
		t.Logf("content %v, attempts %v", string(msg.Content), msg.Attempts)
		count++
		return nil
	})
	if err != nil {
		t.Errorf("subscribe %v", err)
		return
	}
	afterTime := time.Now()
	totalTime := afterTime.Sub(beforeTime)
	wg.Wait()

	t.Logf("messages processed %v, total time %v", count, totalTime)
	t.Log("testsubnsq success exit")
}

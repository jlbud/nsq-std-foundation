package pub

import (
	"fmt"
	"testing"
	"time"

	"github.com/klbud/nsq-std-foundation"
)

func TestPubNSQ(t *testing.T) {
	cfg := map[string]interface{}{
		"lookupds_address": []string{"192.168.1.174:4161"},
	}
	pub, err := New(mq.NSQ, func(log string) error {
		fmt.Println("err log ", log)
		return nil
	}, cfg)
	if err != nil {
		t.Errorf("new err %v", err)
		return
	}

	err = pub.Init("test")
	if err != nil {
		t.Errorf("init err %v", err)
		return
	}

	quit := false
	count := 0
	go func() {
		time.Sleep(10 * time.Second)
		quit = true
	}()

	hello := `
{
 "user_id": 2000032,
 "product_id": 1
}`
	for {
		err = pub.Publish(&mq.Message{
			Content: []byte(hello),
		})
		if err != nil {
			fmt.Println("publish err", err)
		}

		time.Sleep(1 * time.Second)
		count++
		if quit {
			break
		}
	}

	t.Logf("messages pushed %v", count)
	t.Log("testpubnsq success exit")
}

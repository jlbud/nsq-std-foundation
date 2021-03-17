package nsqkit

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"

	"github.com/klbud/nsq-std-foundation/message"
	"github.com/pkg/errors"
)

func TestConsumerServer_StartHandler(t *testing.T) {
	nc := NewConsumer(&ConsumerP{
		TopicName:   "test",
		ChannelName: "a",
		ConsumerConfigP: &ConsumerConfigP{
			DefaultRequeueDelay: 1 * time.Second,
			MsgTimeout:          5 * time.Second,
			MaxAttempts:         10,
			MaxInFlight:         2500,
		},
	})

	wg := sync.WaitGroup{}
	wg.Add(1)

	count := 0
	var msgId [16]byte

	err := nc.StartHandler(message.HandlerFunc(func(message *message.Message) error {
		t.Log("TestConsumerServer_StartHandler Body: ", string(message.Body()), message.Attempts())
		count ++
		if message.ID() == msgId {
			t.Log("the data is repeatedly consumed,attempts ", message.Attempts())
		}
		if count == 1 {
			t.Log("message ID is", message.ID())
			message.Requeue(1 * time.Second)
			msgId = message.ID() //[48 97 98 55 98 99 100 50 102 48 50 57 51 48 48 48]
			return errors.New("test err")
		} else {
			return nil
		}
	}))
	if err != nil {
		t.Log(err.Error())
		wg.Done()
	}
	go func() {
		time.Sleep(2 * time.Second)
		wg.Done()
	}()
	wg.Wait()
	nc.StopConsumer()
	t.Log("test consumer exit")
}

func TestConsumerServer_StartConcurrentHandlers(t *testing.T) {
	nc := NewConsumer(&ConsumerP{
		TopicName:   "test",
		ChannelName: "a",
	})

	wg := sync.WaitGroup{}
	wg.Add(1)

	err := nc.StartConcurrentHandlers(message.HandlerFunc(func(message *message.Message) error {
		message.Finish()
		//do something
		log.Printf("StartConcurrentHandlers: %v ,Attempts: %v", string(message.Body()), message.Attempts())
		return nil
	}))
	if err != nil {
		fmt.Println(err.Error())
		wg.Done()
	}
	go func() {
		time.Sleep(10 * time.Second)
		wg.Done()
	}()
	wg.Wait()
	nc.StopConsumer()
	t.Log("test consumer exit")
}

package pub

import (
	"fmt"
	"github.com/klbud/nsq-std-foundation"
	nsq "github.com/klbud/nsq-std-foundation/nsq/defaultmq"
)

// create mq-pub ins, include config check
// success created means connecting TCP or something success
func New(mqType mq.MQType, errorLog func(log string) error, config map[string]interface{}) (mq.Pub, error) {
	switch mqType {
	case mq.NSQ:
		return nsq.NewPub(errorLog, config)
	}
	return nil, fmt.Errorf("not achieve")
}

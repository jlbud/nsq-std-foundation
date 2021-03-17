package sub

import (
	"fmt"
	"github.com/klbud/nsq-std-foundation"
	nsq "github.com/klbud/nsq-std-foundation/nsq/defaultmq"
)

// create mq-sub ins, include config check
// success created means connecting TCP or something success
func New(mqType mq.MQType, errorLog func(log string) error, config map[string]interface{}) (mq.Sub, error) {
	switch mqType {
	case mq.NSQ:
		return nsq.NewSub(errorLog, config)
	}
	return nil, fmt.Errorf("not achieve")
}

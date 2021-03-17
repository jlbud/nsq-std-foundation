package config

import (
	"github.com/klbud/nsq-std-foundation/mock"
)

var (
	LookupdsAddress []string //lookupds对业务接口地址
)

func init() {
	LookupdsAddress = mock.LookupdHost
}

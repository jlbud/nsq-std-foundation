package config

import (
	"github.com/Kevin005/nsq-std-foundation/mock"
)

var (
	LookupdsAddress []string //lookupds对业务接口地址
)

func init() {
	LookupdsAddress = mock.LookupdHost
}

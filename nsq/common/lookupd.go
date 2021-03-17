package common

import (
	"encoding/json"
	"errors"

	"github.com/klbud/nsq-std-foundation/nsq/util/http"
)

type Lookupd struct {
	LookupdAddress []string // 服务发现服务地址列表
}

// 从服务发现服务中获取可用的nsqd
func (l *Lookupd) GetNsqds() (error, *Nsqds) {
	if l.LookupdAddress == nil || len(l.LookupdAddress) == 0 {
		return errors.New("getnsqds err, lookupdaddress empty"), nil
	}

	for _, lkaddress := range l.LookupdAddress {
		err, nds := l.getNsqds(lkaddress)
		if err != nil {
			continue
		}
		if len(nds.Producers) > 0 {
			return nil, nds
		}
	}
	return errors.New("no nsqds was found available"), nil
}

type Nsqds struct {
	Producers []Producer `json:"producers"`
}

type Producer struct {
	HttpPort         int    `json:"Http_port"`
	TcpPort          int    `json:"tcp_port"`
	RemoteAddress    string `json:"remote_address"`
	BroadcastAddress string `json:"broadcast_address"`
}

func (l *Lookupd) getNsqds(lkaddress string) (error, *Nsqds) {
	hr := &http.HttpResult{}
	err := http.SendGet(lkaddress+GetAllNsqdsUrl(), hr)
	if err != nil {
		return err, nil
	}
	if hr.StateCode == http.HttpResultSuccess {
		ns := &Nsqds{}
		json.Unmarshal([]byte(hr.Body), ns)
		return nil, ns
	} else {
		return errors.New(hr.Body), nil
	}
}

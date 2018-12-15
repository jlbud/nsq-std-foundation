package mq

import (
	"testing"
)

func TestLookupd_GetNsqd(t *testing.T) {
	ld := &Lookupd{
		LookupdAddress: []string{"http://127.0.1.174:4161"},
	}
	err, ns := ld.GetNsqds()
	if err != nil {
		t.Error(err.Error())
		return
	}
	for i, n := range ns.Producers {
		t.Log(i, n.HttpPort)
		t.Log(i, n.BroadcastAddress)
	}
}

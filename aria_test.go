package aria2rpc_test

import (
	"testing"

	"github.com/xzzpig/aria2rpc"
)

var aria = aria2rpc.Aria{
	RpcUrl: "http://token:52xzzpig@127.0.0.1:6800/jsonrpc",
	Secret: "52xzzpig",
}

func TestPauseAll(t *testing.T) {
	err := aria.PauseAll()
	if err != nil {
		t.Error(err)
	}
}

func TestGetVersion(t *testing.T) {
	version, err := aria.GetVersion()
	if err != nil {
		t.Error(err)
	}
	t.Log(version)
}

func TestListMethods(t *testing.T) {
	result, err := aria.ListMethods()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestListNotifications(t *testing.T) {
	result, err := aria.ListNotifications()
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestRemove(t *testing.T) {
	result, err := aria.Remove("1")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestTellActive(t *testing.T) {
	result, err := aria.TellActive(nil)
	if err != nil {
		t.Error(err)
	}

	t.Log(result)
	for _, status := range result {
		t.Log(status.Gid)
	}
}

func TestTellStatus(t *testing.T) {
	result, err := aria.TellStatus("6718ec25e4b45254", nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

func TestGetFiles(t *testing.T) {
	result, err := aria.GetFiles("6718ec25e4b45254")
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

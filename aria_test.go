package aria2rpc_test

import (
	"testing"

	"github.com/xzzpig/aria2rpc"
	"github.com/xzzpig/aria2rpc/types"
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

func TestAddUri(t *testing.T) {
	result, err := aria.AddUri([]string{"https://dm.phncdn.com/videos/202010/21/362574322/1080P_8000K_362574322.mp4?ttl=1617243704&ri=1638400&rs=4000&hash=663bc32c69135a49c3dedfe4cda841d8"}, types.Aria2Options{})
	if err != nil {
		t.Error(err)
	}
	t.Log(result)
}

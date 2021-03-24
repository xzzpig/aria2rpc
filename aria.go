package aria2rpc

import (
	"strings"

	"github.com/kdar/httprpc"
	"github.com/xzzpig/aria2rpc/types"
)

type Aria struct {
	RpcUrl string
	Secret string
}

var RpcVersion = "2.0"

func prefix(str string) string {
	if !strings.HasPrefix(str, "system.") && !strings.HasPrefix(str, "aria2.") {
		str = "aria2." + str
	}
	return str
}

func (aria *Aria) addSecret(params []interface{}) []interface{} {
	if aria.Secret == "" {
		return params
	}
	p := make([]interface{}, 0)
	p = append(p, "token:"+aria.Secret)
	p = append(p, params...)
	return p
}

func (aria *Aria) Call(method string, params []interface{}, reply interface{}) error {
	method = prefix(method)
	params = aria.addSecret(params)
	err := httprpc.CallJson(RpcVersion, aria.RpcUrl, method, params, reply)
	if err != nil {
		return err
	}
	return nil
}

func (aria *Aria) ListMethods() (result []string, err error) {
	err = aria.Call("system.listMethods", make([]interface{}, 0), &result)
	return
}

func (aria *Aria) ListNotifications() (result []string, err error) {
	err = aria.Call("system.listNotifications", make([]interface{}, 0), &result)
	return
}

func (aria *Aria) AddUri(uris []string, options types.Aria2Options) (result types.GID, err error) {
	err = aria.Call("addUri", append(make([]interface{}, 0), uris, options), &result)
	return
}

func (aria *Aria) Remove(gid types.GID) (result types.GID, err error) {
	err = aria.Call("remove", append(make([]interface{}, 0), gid), &result)
	return
}

func (aria *Aria) ForceRemove(gid types.GID) (result types.GID, err error) {
	err = aria.Call("forceRemove", append(make([]interface{}, 0), gid), &result)
	return
}

func (aria *Aria) Pause(gid types.GID) (result types.GID, err error) {
	err = aria.Call("pause", append(make([]interface{}, 0), gid), &result)
	return
}

func (aria *Aria) ForcePause(gid types.GID) (result types.GID, err error) {
	err = aria.Call("forcePause", append(make([]interface{}, 0), gid), &result)
	return
}

func (aria *Aria) PauseAll() error {
	var reply interface{}
	err := aria.Call("pauseAll", make([]interface{}, 0), &reply)
	if err != nil {
		return err
	}
	return nil
}

func (aria *Aria) ForcePauseAll() error {
	var reply interface{}
	err := aria.Call("forcePauseAll", make([]interface{}, 0), &reply)
	if err != nil {
		return err
	}
	return nil
}

func (aria *Aria) Unpause(gid types.GID) (result types.GID, err error) {
	err = aria.Call("unpause", append(make([]interface{}, 0), gid), &result)
	return
}

func (aria *Aria) UnpauseAll() error {
	var reply interface{}
	err := aria.Call("unpauseAll", make([]interface{}, 0), &reply)
	if err != nil {
		return err
	}
	return nil
}

// keys: keyof DownloadStatus
func (aria *Aria) TellStatus(gid string, keys []string) (result types.DownloadStatus, err error) {
	if len(keys) == 0 {
		err = aria.Call("tellStatus", append(make([]interface{}, 0), gid), &result)
	} else {
		err = aria.Call("tellStatus", append(make([]interface{}, 0), gid, keys), &result)
	}
	return
}

func (aria *Aria) GetUris(gid types.GID) (result []types.AriaUri, err error) {
	err = aria.Call("getUris", append(make([]interface{}, 0), gid), &result)
	return
}

func (aria *Aria) GetFiles(gid types.GID) (result []types.AriaFile, err error) {
	err = aria.Call("getFiles", append(make([]interface{}, 0), gid), &result)
	return
}

// keys: keyof DownloadStatus
func (aria *Aria) TellActive(keys []string) (result []types.DownloadStatus, err error) {
	if len(keys) == 0 {
		err = aria.Call("tellActive", make([]interface{}, 0), &result)
	} else {
		err = aria.Call("tellActive", append(make([]interface{}, 0), keys), &result)
	}
	return
}

// keys: keyof DownloadStatus
func (aria *Aria) TellWaiting(offset int, num int, keys []string) (result []types.DownloadStatus, err error) {
	if len(keys) == 0 {
		err = aria.Call("tellWaiting", append(make([]interface{}, 0), offset, num), &result)
	} else {
		err = aria.Call("tellWaiting", append(make([]interface{}, 0), offset, num, keys), &result)
	}
	return
}

// keys: keyof DownloadStatus
func (aria *Aria) TellStopped(offset int, num int, keys []string) (result []types.DownloadStatus, err error) {
	if len(keys) == 0 {
		err = aria.Call("tellStopped", append(make([]interface{}, 0), offset, num), &result)
	} else {
		err = aria.Call("tellStopped", append(make([]interface{}, 0), offset, num, keys), &result)
	}
	return
}

func (aria *Aria) GetGlobalStat() (result types.GlobalStat, err error) {
	err = aria.Call("getGlobalStat", make([]interface{}, 0), &result)
	return
}

func (aria *Aria) PurgeDownloadResult() (result interface{}, err error) {
	err = aria.Call("purgeDownloadResult", make([]interface{}, 0), &result)
	return
}

func (aria *Aria) RemoveDownloadResult(gid types.GID) (result types.GID, err error) {
	err = aria.Call("removeDownloadResult", append(make([]interface{}, 0), gid), &result)
	return
}

func (aria *Aria) GetVersion() (result types.AriaVersion, err error) {
	err = aria.Call("getVersion", make([]interface{}, 0), &result)
	return
}

func (aria *Aria) Shutdown() (result interface{}, err error) {
	err = aria.Call("shutdown", make([]interface{}, 0), &result)
	return
}

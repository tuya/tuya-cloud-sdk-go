package infrared

import (
	"encoding/json"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PutSetAliasReq struct {
	InfraredID string
	RemoteID   string `json:"remote_id"`
	RemoteName string `json:"remote_name"`
}

func (t *PutSetAliasReq) Method() string {
	return common.RequestPut
}

func (t *PutSetAliasReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s", t.InfraredID)
}

func (t *PutSetAliasReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PutSetAlias sets infrared learning mode
func PutSetAlias(infraredID, remoteID, remoteName string) (*PutSetAliasResponse, error) {
	a := &PutSetAliasReq{InfraredID: infraredID, RemoteID: remoteID, RemoteName: remoteName}
	resp := &PutSetAliasResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type PutSetAliasResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

package infrared

import (
	"encoding/json"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PostCommandReq struct {
	InfraredID string
	RemoteID   int
	Key        string `json:"key"`
}

func (t *PostCommandReq) Method() string {
	return common.RequestPost
}

func (t *PostCommandReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/remotes/%s/command", t.InfraredID, t.RemoteID)
}

func (t *PostCommandReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PostCommand sends command to infrared controlled device
func PostCommand(infraredID string, remoteId int, key string) (*PostCommandResponse, error) {
	a := &PostCommandReq{
		InfraredID: infraredID,
		RemoteID:   remoteId,
		Key:        key,
	}
	resp := &PostCommandResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type PostCommandResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

package infrared

import (
	"encoding/json"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PostTestingCommandReq struct {
	InfraredID  string
	RemoteIndex int    `json:"remote_index"`
	CategoryId  int    `json:"category_id"`
	Key         string `json:"key"`
}

func (t *PostTestingCommandReq) Method() string {
	return common.RequestPost
}

func (t *PostTestingCommandReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/testing/command", t.InfraredID)
}

func (t *PostTestingCommandReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PostTestingCommand sends test command to infrared controlled device
func PostTestingCommand(infraredID string, remoteIndex, categoryId int, key string) (*PostTestingCommandResponse, error) {
	a := &PostTestingCommandReq{
		InfraredID:  infraredID,
		CategoryId:  categoryId,
		RemoteIndex: remoteIndex,
		Key:         key,
	}
	resp := &PostTestingCommandResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type PostTestingCommandResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

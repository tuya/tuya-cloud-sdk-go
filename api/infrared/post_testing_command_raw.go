package infrared

import (
	"encoding/json"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PostTestingCommandRawReq struct {
	InfraredID  string
	RemoteIndex int `json:"remote_index"`
	CategoryId  int `json:"category_id"`
	RawKey      int `json:"raw_key"`
}

func (t *PostTestingCommandRawReq) Method() string {
	return common.RequestPost
}

func (t *PostTestingCommandRawReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/testing/raw/command", t.InfraredID)
}

func (t *PostTestingCommandRawReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PostTestingCommandRaw sends test command to infrared controlled device
func PostTestingCommandRaw(infraredID string, remoteIndex, categoryId int, rawKey int) (*PostTestingCommandRawResponse, error) {
	a := &PostTestingCommandRawReq{
		InfraredID:  infraredID,
		CategoryId:  categoryId,
		RemoteIndex: remoteIndex,
		RawKey:      rawKey,
	}
	resp := &PostTestingCommandRawResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type PostTestingCommandRawResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

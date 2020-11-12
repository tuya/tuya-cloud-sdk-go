package infrared

import (
	"encoding/json"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PostCommandRawReq struct {
	InfraredID string
	RemoteID   int
	RawKey     int `json:"raw_key"`
}

func (t *PostCommandRawReq) Method() string {
	return common.RequestPost
}

func (t *PostCommandRawReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/remotes/%s/raw/command", t.InfraredID, t.RemoteID)
}

func (t *PostCommandRawReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PostCommandRaw sends command to infrared controlled device based on pairing rules
func PostCommandRaw(infraredID string, remoteId int, rawKey int) (*PostCommandRawResponse, error) {
	a := &PostCommandRawReq{
		InfraredID: infraredID,
		RemoteID:   remoteId,
		RawKey:     rawKey,
	}
	resp := &PostCommandRawResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type PostCommandRawResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

package infrared

import (
	"encoding/json"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PostACCommandReq struct {
	InfraredID string
	RemoteId   string
	Code       string `json:"code"`  // one of "power", "mode", "temp", or "wind"
	Value      int    `json:"value"` // See https://developer.tuya.com/en/docs/iot/open-api/api-list/api/proprietary-category/universal-infrared?id=K9jgsgd7buln4#title-29-Air%C2%A0conditioning%C2%A0remote%C2%A0control
}

func (t *PostACCommandReq) Method() string {
	return common.RequestPost
}

func (t *PostACCommandReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/air-conditioners/%s/command", t.InfraredID, t.RemoteId)
}

func (t *PostACCommandReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PostACCommand sends test command to infrared controlled Air Conditioner.
// See https://developer.tuya.com/en/docs/iot/open-api/api-list/api/proprietary-category/universal-infrared?id=K9jgsgd7buln4#title-29-Air%C2%A0conditioning%C2%A0remote%C2%A0control
// for details
func PostACCommand(infraredID, remoteId, code string, value int) (*PostACCommandResponse, error) {
	a := &PostACCommandReq{
		InfraredID: infraredID,
		RemoteId:   remoteId,
		Code:       code,
		Value:      value,
	}
	resp := &PostACCommandResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type PostACCommandResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

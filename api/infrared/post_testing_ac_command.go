package infrared

import (
	"encoding/json"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PostTestingACCommandReq struct {
	InfraredID  string
	RemoteIndex int    `json:"remote_index"`
	CategoryId  int    `json:"category_id"`
	Code        string `json:"code"`  // one of "power", "mode", "temp", or "wind"
	Value       int    `json:"value"` // See https://developer.tuya.com/en/docs/iot/open-api/api-list/api/proprietary-category/universal-infrared?id=K9jgsgd7buln4#title-28-Test%C2%A0air%C2%A0conditioner%C2%A0remote%C2%A0control
}

func (t *PostTestingACCommandReq) Method() string {
	return common.RequestPost
}

func (t *PostTestingACCommandReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/air-conditioners/testing/command", t.InfraredID)
}

func (t *PostTestingACCommandReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PostTestingACCommand sends test command to infrared controlled Air Conditioner.
// See https://developer.tuya.com/en/docs/iot/open-api/api-list/api/proprietary-category/universal-infrared?id=K9jgsgd7buln4#title-28-Test%C2%A0air%C2%A0conditioner%C2%A0remote%C2%A0control
// for details
func PostTestingACCommand(infraredID string, remoteIndex, categoryId int, code string, value int) (*PostTestingACCommandResponse, error) {
	a := &PostTestingACCommandReq{
		InfraredID:  infraredID,
		CategoryId:  categoryId,
		RemoteIndex: remoteIndex,
		Code:        code,
		Value:       value,
	}
	resp := &PostTestingACCommandResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type PostTestingACCommandResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

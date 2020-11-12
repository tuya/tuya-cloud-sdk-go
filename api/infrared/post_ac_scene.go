package infrared

import (
	"encoding/json"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PostACSceneReq struct {
	InfraredID string
	RemoteId   string
	Power      int `json:"power"`
	Mode       int `json:"mode"`
	Temp       int `json:"temp"`
	Wind       int `json:"wind"`
}

func (t *PostACSceneReq) Method() string {
	return common.RequestPost
}

func (t *PostACSceneReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/air-conditioners/%s/scenes/command", t.InfraredID, t.RemoteId)
}

func (t *PostACSceneReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PostACScene sends scene command to infrared controlled Air Conditioner.
// See https://developer.tuya.com/en/docs/iot/open-api/api-list/api/proprietary-category/universal-infrared?id=K9jgsgd7buln4#title-30-Multi-condition%C2%A0control%C2%A0air%C2%A0conditioner
// for details
func PostACScene(infraredID, remoteId string, power, mode, temp, wind int) (*PostACSceneResponse, error) {
	a := &PostACSceneReq{
		InfraredID: infraredID,
		RemoteId:   remoteId,
		Power:      power,
		Mode:       mode,
		Temp:       temp,
		Wind:       wind,
	}
	resp := &PostACSceneResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type PostACSceneResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

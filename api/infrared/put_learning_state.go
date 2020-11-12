package infrared

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PutLearningStateReq struct {
	InfraredID string
	State      bool
}

func (t *PutLearningStateReq) Method() string {
	return common.RequestPut
}

func (t *PutLearningStateReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/learning-state?state=%t", t.InfraredID, t.State)
}

// PutLearningState sets infrared learning mode
func PutLearningState(infraredID string, state bool) (*PutLearningStateResponse, error) {
	a := &PutLearningStateReq{InfraredID: infraredID, State: state}
	resp := &PutLearningStateResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type PutLearningStateResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

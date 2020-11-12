package infrared

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetLearningCodesReq struct {
	InfraredID string
	Time       int64
}

func (t *GetLearningCodesReq) Method() string {
	return common.RequestGet
}

func (t *GetLearningCodesReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/learning-codes?learning_time=%d", t.InfraredID, t.Time)
}

// GetLearningCodes sets infrared learning mode
func GetLearningCodes(infraredID string, time int64) (*GetLearningCodesResponse, error) {
	a := &GetLearningCodesReq{InfraredID: infraredID, Time: time}
	resp := &GetLearningCodesResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetLearningCodesResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		Success bool   `json:"success"`
		Code    string `json:"code"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

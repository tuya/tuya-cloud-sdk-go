package device_dn

import "fmt"

import "github.com/TuyaInc/tuya_cloud_sdk_go/api/common"

type GetDevicesByTokenReq struct {
	PairToken string
}

func (t *GetDevicesByTokenReq) Method() string {
	return common.RequestGet
}

func (t *GetDevicesByTokenReq) API() string {
	return fmt.Sprintf("/v1.0/devices/tokens/%s", t.PairToken)
}

func GetDevicesByToken(pairToken string) (*GetDevicesByTokenResponse, error) {
	req := &GetDevicesByTokenReq{
		PairToken: pairToken,
	}
	resp := &GetDevicesByTokenResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type GetDevicesByTokenResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		SuccessDevices []interface{} `json:"successDevices"`
		ErrorDevices   []interface{} `json:"errorDevices"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

package device

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetDeviceSpecReq struct {
	DeviceID string
}

func (t *GetDeviceSpecReq) Method() string {
	return common.RequestGet
}

func (t *GetDeviceSpecReq) API() string {
	return fmt.Sprintf("/v1.0/devices/%s/specifications", t.DeviceID)
}

// GetDeviceFunction Get the function list based on the device id
func GetDeviceSpecifications(deviceID string) (*GetDeviceSpecResponse, error) {
	a := &GetDeviceSpecReq{DeviceID: deviceID}
	resp := &GetDeviceSpecResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetDeviceSpecResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		Category  string `json:"category"`
		Functions []struct {
			Code   string `json:"code"`
			Type   string `json:"type"`
			Values string `json:"values"`
		} `json:"functions"`
		Status []struct {
			Code   string `json:"code"`
			Type   string `json:"type"`
			Values string `json:"values"`
		} `json:"status"`
	} `json:"result"`

	//  error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

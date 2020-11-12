package device

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetDeviceFunctionReq struct {
	DeviceID string
}

func (t *GetDeviceFunctionReq) Method() string {
	return common.RequestGet
}

func (t *GetDeviceFunctionReq) API() string {
	return fmt.Sprintf("/v1.0/devices/%s/functions", t.DeviceID)
}

// GetDeviceFunction Get the function list based on the device id
func GetDeviceFunctions(deviceID string) (*GetDeviceFunctionResponse, error) {
	a := &GetDeviceFunctionReq{DeviceID: deviceID}
	resp := &GetDeviceFunctionResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetDeviceFunctionResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		Category  string `json:"category"`
		Functions []struct {
			Name   string `json:"name"`
			Desc   string `json:"desc"`
			Code   string `json:"code"`
			Type   string `json:"type"`
			Values string `json:"values"`
		} `json:"functions"`
	} `json:"result"`

	//  error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

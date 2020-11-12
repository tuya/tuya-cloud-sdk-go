package device

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetDeviceStatusReq struct {
	DeviceID string
}

func (t *GetDeviceStatusReq) Method() string {
	return common.RequestGet
}

func (t *GetDeviceStatusReq) API() string {
	return fmt.Sprintf("/v1.0/devices/%s/status", t.DeviceID)
}

// GetDeviceStatus Obtain device function point messages based on device ID
func GetDeviceStatus(deviceID string) (*GetDeviceStatusResponse, error) {
	a := &GetDeviceStatusReq{DeviceID: deviceID}
	resp := &GetDeviceStatusResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetDeviceStatusResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  []struct {
		Code  string      `json:"code"`
		Value interface{} `json:"value"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

package device

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type DeleteDeviceReq struct {
	DeviceID string
}

func (t *DeleteDeviceReq) Method() string {
	return common.RequestDelete
}

func (t *DeleteDeviceReq) API() string {
	return fmt.Sprintf("/v1.0/devices/%s", t.DeviceID)
}

// DeleteDevice 移除设备
func DeleteDevice(deviceID string) (*DeleteDeviceResponse, error) {
	a := &DeleteDeviceReq{DeviceID: deviceID}
	resp := &DeleteDeviceResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type DeleteDeviceResponse struct {
	Success bool  `json:"success"`
	Result  bool  `json:"result"`
	T       int64 `json:"t"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

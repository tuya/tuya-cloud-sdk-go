package device

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetDeviceReq struct {
	DeviceID string
}

func (t *GetDeviceReq) Method() string {
	return common.RequestGet
}

func (t *GetDeviceReq) API() string {
	return fmt.Sprintf("/v1.0/devices/%s", t.DeviceID)
}

// GetDevice 获取设备信息
func GetDevice(deviceID string) (*GetDeviceResponse, error) {
	a := &GetDeviceReq{DeviceID: deviceID}
	resp := &GetDeviceResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetDeviceResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		UUID   string `json:"uuid"`
		UID    string `json:"uid"`
		Name   string `json:"name"`
		IP     string `json:"ip"`
		Sub    bool   `json:"sub"`
		Model  string `json:"model"`
		Status []struct {
			Code  string      `json:"code"`
			Value interface{} `json:"value"`
		} `json:"status"`
		Category    string `json:"category"`
		Online      bool   `json:"online"`
		ID          string `json:"id"`
		TimeZone    string `json:"time_zone"`
		LocalKey    string `json:"local_key"`
		UpdateTime  int    `json:"update_time"`
		ActiveTime  int    `json:"active_time"`
		OwnerID     string `json:"owner_id"`
		ProductID   string `json:"product_id"`
		ProductName string `json:"product_name"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

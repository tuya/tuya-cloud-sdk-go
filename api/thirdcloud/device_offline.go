package thirdcloud

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PutDevicesOfflineReq struct {
	ThirdCloudDeviceId string // 第三⽅方云设备id
}

func (t *PutDevicesOfflineReq) Method() string {
	return common.RequestPut
}

func (t *PutDevicesOfflineReq) API() string {
	// /v1.0/3rdcloud/devices/{3rd_cloud_device_id}/offline
	return fmt.Sprintf("/v1.0/3rdcloud/devices/%s/offline", t.ThirdCloudDeviceId)
}

func PutDevicesOffline(thirdCloudDeviceId string) (*PutDevicesOfflineResponse, error) {
	req := &PutDevicesOfflineReq{
		ThirdCloudDeviceId: thirdCloudDeviceId,
	}
	resp := &PutDevicesOfflineResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type PutDevicesOfflineResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

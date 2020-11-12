package thirdcloud

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PutDevicesOnlineReq struct {
	ThirdCloudDeviceId string // 第三⽅方云设备id
}

func (t *PutDevicesOnlineReq) Method() string {
	return common.RequestPut
}

func (t *PutDevicesOnlineReq) API() string {
	// /v1.0/3rdcloud/devices/{3rd_cloud_device_id}/online
	return fmt.Sprintf("/v1.0/3rdcloud/devices/%s/online", t.ThirdCloudDeviceId)
}

func PutDevicesOnline(thirdCloudDeviceId string) (*PutDevicesOnlineResponse, error) {
	req := &PutDevicesOnlineReq{
		ThirdCloudDeviceId: thirdCloudDeviceId,
	}
	resp := &PutDevicesOnlineResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type PutDevicesOnlineResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

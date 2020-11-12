package thirdcloud

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type DeleteDevicesUnBindReq struct {
	ThirdCloudDeviceId string // 第三⽅方云设备id
}

func (t *DeleteDevicesUnBindReq) Method() string {
	return common.RequestDelete
}

func (t *DeleteDevicesUnBindReq) API() string {
	// /v1.0/3rdcloud/devices/{3rd_cloud_device_id}/unbind
	return fmt.Sprintf("/v1.0/3rdcloud/devices/%s/unbind", t.ThirdCloudDeviceId)
}

func DeleteDevicesUnBind(thirdCloudDeviceId string) (*DeleteDevicesUnBindResponse, error) {
	req := &DeleteDevicesUnBindReq{
		ThirdCloudDeviceId: thirdCloudDeviceId,
	}
	resp := &DeleteDevicesUnBindResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type DeleteDevicesUnBindResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

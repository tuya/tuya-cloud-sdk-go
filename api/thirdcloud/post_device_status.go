package thirdcloud

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PostDevicesStatusReq struct {
	ThirdCloudDeviceId string                   //	第三⽅方云设备id
	Status             PostDevicesStatusReqBody //	设备状态数据
}

type PostDevicesStatusReqBody struct {
	Status []PostDevicesStatusItem `json:"status"` // 设备状态
}
type PostDevicesStatusItem struct {
	Code  string      `json:"code"`  // 功能点code
	Value interface{} `json:"value"` // 功能点的值
}

func (t *PostDevicesStatusReq) Method() string {
	return common.RequestPost
}

func (t *PostDevicesStatusReq) API() string {
	// /v1.0/3rdcloud/devices/{3rd_cloud_device_id}/status
	return fmt.Sprintf("/v1.0/3rdcloud/devices/%s/status", t.ThirdCloudDeviceId)
}

func (t *PostDevicesStatusReq) Body() []byte {
	bs, _ := json.Marshal(t.Status)
	return bs
}

func PostDevicesStatus(thirdCloudDeviceId string, statusItems []PostDevicesStatusItem) (*PostDevicesStatusResponse, error) {
	if len(statusItems) == 0 {
		return nil, errors.New("statusItem is nil")
	}
	req := &PostDevicesStatusReq{
		ThirdCloudDeviceId: thirdCloudDeviceId,
		Status: PostDevicesStatusReqBody{
			Status: statusItems,
		},
	}
	resp := &PostDevicesStatusResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type PostDevicesStatusResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

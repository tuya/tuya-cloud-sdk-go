package pairing

import (
	"fmt"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

type GetSubDeviceListReq struct {
	DeviceId string
}

func (t *GetSubDeviceListReq) Method() string {
	return common.RequestGet
}

func (t *GetSubDeviceListReq) API() string {
	return fmt.Sprintf("/v1.0/devices/%s/sub-devices", t.DeviceId)
}

// GetSubDeviceList Get the list of sub-devices through the device IDs under the gateway
func GetSubDeviceList(deviceId string) (*GetSubDeviceListResponse, error) {
	req := &GetSubDeviceListReq{DeviceId: deviceId}
	resp := &GetSubDeviceListResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type GetSubDeviceListResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  []struct {
		ID         string `json:"id"`          // ID
		ProductID  string `json:"product_id"`  // Product ID
		OwnerId    string `json:"owner_id"`    // Device owner ID
		Online     bool   `json:"online"`      // Online status of device
		Name       string `json:"name"`        // Device name
		UpdateTime int64  `json:"update_time"` // The update time of device status
		ActiveTime int64  `json:"active_time"` // The last pairing time of the device
		Category   string `json:"category"`    // Category
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

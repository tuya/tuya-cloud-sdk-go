package pairing

import (
	"fmt"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

type EnableSubDiscoveryReq struct {
	DeviceId string `json:"device_id"` // Gateway device id
	Duration int    `json:"duration"`  // Gateway
}

func (t *EnableSubDiscoveryReq) Method() string {
	return common.RequestPut
}

func (t *EnableSubDiscoveryReq) API() string {
	return fmt.Sprintf("/v1.0/devices/%s/enabled-sub-discovery?duration=%d", t.DeviceId, t.Duration)
}

// EnableSubDiscovery Obtain pairing information
func EnableSubDiscovery(deviceId string, duration int) (*EnableSubDiscoveryResponse, error) {
	req := &EnableSubDiscoveryReq{
		DeviceId: deviceId,
		Duration: duration,
	}
	resp := &EnableSubDiscoveryResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type EnableSubDiscoveryResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

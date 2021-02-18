package pairing

import (
	"fmt"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

type GetPairingDevicesReq struct {
	Token string
}

func (t *GetPairingDevicesReq) Method() string {
	return common.RequestGet
}

func (t *GetPairingDevicesReq) API() string {
	return fmt.Sprintf("/v1.0/device/paring/tokens/%s", t.Token)
}

// GetPairingDevices Obtain App users based on schema page
func GetPairingDevices(token string) (*GetPairingDevicesResponse, error) {
	req := &GetPairingDevicesReq{Token: token}
	resp := &GetPairingDevicesResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type GetPairingDevicesResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		Success []Success `json:"success"` // List of devices whose pairing succeeded
		Failed  []Failed  `json:"failed"`  // List of devices whose pairing failed
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type Success struct {
	DeviceID  string `json:"device_id"`  // IDs of devices whose pairing succeeded
	ProductID string `json:"product_id"` // IDs of products whose pairing succeeded
	Name      string `json:"name"`       // Names of devices whose pairing succeeded
	Category  string `json:"category"`   // Types of devices whose pairing succeeded
}

type Failed struct {
	DeviceID string `json:"device_id"` // IDs of devices whose pairing failed
	Code     string `json:"code"`      // Failure error code
	Msg      string `json:"msg"`       // Description of failure error
	Name     string `json:"name"`      // Names of devices whose pairing failed
}

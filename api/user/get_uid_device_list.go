package user

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetDeviceListByUIDReq struct {
	UID string
}

func (t *GetDeviceListByUIDReq) Method() string {
	return common.RequestGet
}

func (t *GetDeviceListByUIDReq) API() string {
	return fmt.Sprintf("/v1.0/users/%s/devices", t.UID)
}

func GetDeviceListByUID(uid string) (*GetDeviceListByUIDResponse, error) {
	req := &GetDeviceListByUIDReq{UID: uid}
	resp := &GetDeviceListByUIDResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type GetDeviceListByUIDResponse struct {
	Success bool          `json:"success"`
	T       int64         `json:"t"`
	Result  []interface{} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

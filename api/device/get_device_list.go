package device

import (
	"fmt"
	"strings"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetDeviceListReq struct {
	DeviceIDList []string
	PageNo       uint
	PageSize     uint
}

func (t *GetDeviceListReq) Method() string {
	return common.RequestGet
}

func (t *GetDeviceListReq) API() string {
	if t.PageNo <= 0 {
		t.PageNo = 1
	}
	if t.PageSize <= 0 {
		t.PageSize = 20
	}
	return fmt.Sprintf(
		"/v1.0/devices?device_ids=%s&page_no=%d&page_size=%d",
		strings.Join(t.DeviceIDList, ","),
		t.PageNo,
		t.PageSize)
}

func GetDeviceList(deviceIDList []string) (*GetDeviceListResponse, error) {
	a := &GetDeviceListReq{DeviceIDList: deviceIDList}
	resp := &GetDeviceListResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetDeviceListResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		Total   int `json:"total"`
		Devices []struct {
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
			Category   string `json:"category"`
			Online     bool   `json:"online"`
			ID         string `json:"id"`
			TimeZone   string `json:"time_zone"`
			LocalKey   string `json:"local_key"`
			UpdateTime int    `json:"update_time"`
			ActiveTime int    `json:"active_time"`
			OwnerID    string `json:"owner_id"`
			ProductID  string `json:"product_id"`
		} `json:"devices"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

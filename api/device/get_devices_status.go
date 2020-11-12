package device

import (
	"fmt"
	"log"
	"strings"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetDevicesStatusReq struct {
	DeviceIDList []string
	PageNo       uint
	PageSize     uint
}

func (t *GetDevicesStatusReq) Method() string {
	return common.RequestGet
}

func (t *GetDevicesStatusReq) API() string {
	if t.PageNo <= 0 {
		t.PageNo = 1
	}
	if t.PageSize <= 0 {
		t.PageSize = 20
	}
	a := fmt.Sprintf(
		"/v1.0/devices/status?device_ids=%s&page_no=%d&page_size=%d",
		strings.Join(t.DeviceIDList, ","),
		t.PageNo,
		t.PageSize)
	log.Println("a:", a)
	return a
}

func GetDevicesStatus(deviceIDList []string) (*GetDevicesStatusResponse, error) {
	a := &GetDevicesStatusReq{DeviceIDList: deviceIDList}
	resp := &GetDevicesStatusResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetDevicesStatusResponse struct {
	Result map[string][]struct {
		Code  string      `json:"code"`
		Value interface{} `json:"value"`
	} `json:"result"`
	Success bool  `json:"success"`
	T       int64 `json:"t"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

package user

import (
	"fmt"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
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
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  []struct {
		ID         string `json:"id"`
		UID        string `json:"uid"`
		UUID       string `json:"uuid"`
		LocalKey   string `json:"local_key"`
		Category   string `json:"category"`
		ProductID  string `json:"product_id"`
		Sub        bool   `json:"sub"`
		OwnerID    string `json:"owner_id"`
		Online     bool   `json:"online"`
		Name       string `json:"name"`
		IP         string `json:"ip"`
		TimeZone   string `json:"time_zone"`
		CreateTime int64  `json:"create_time"`
		UpdateTime int64  `json:"update_time"`
		ActiveTime int64  `json:"active_time"`
		Status     []struct {
			Code  string      `json:"code"`
			Value interface{} `json:"value"`
		} `json:"status"`
	} `json:"result"`
	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

package thirdcloud

import (
	"encoding/json"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PostDevicesBindReq struct {
	ThirdCloudDeviceId string // 第三⽅方云设备id
	ReqBody            PostDevicesBindReqBody
}

type PostDevicesBindReqBody struct {
	CountryCode    string `json:"country_code"`     // 国家码
	AppSchema      string `json:"app_schema"`       // 涂鸦应⽤用标识
	TuyaUsername   string `json:"tuya_username"`    // 涂鸦⽤用户名(对应第三⽅方⽤用户唯⼀一标识)
	TuyaProductId  string `json:"tuya_product_id"`  // 涂鸦产品id
	ParentDeviceId string `json:"parent_device_id"` // 第三⽅方⽗父设备id，单品则不不填
}

func (t *PostDevicesBindReq) Method() string {
	return common.RequestPost
}

func (t *PostDevicesBindReq) API() string {
	// /v1.0/3rdcloud/devices/{3rd_cloud_device_id}/bind
	return fmt.Sprintf("/v1.0/3rdcloud/devices/%s/bind", t.ThirdCloudDeviceId)
}

func (t *PostDevicesBindReq) Body() []byte {
	bs, _ := json.Marshal(t.ReqBody)
	return bs
}

func PostDevicesBind(thirdCloudDeviceId, countryCode, appSchema, tuyaUsername, tuyaProductId, parentDeviceId string) (*PostDevicesBindResponse, error) {
	rb := PostDevicesBindReqBody{
		CountryCode:   countryCode,
		AppSchema:     appSchema,
		TuyaUsername:  tuyaUsername,
		TuyaProductId: tuyaProductId,
	}
	if len(parentDeviceId) == 0 {
		rb.ParentDeviceId = parentDeviceId
	}
	req := &PostDevicesBindReq{
		ThirdCloudDeviceId: thirdCloudDeviceId,
		ReqBody:            rb,
	}
	resp := &PostDevicesBindResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

func PostDevicesBindByReq(req *PostDevicesBindReq) (*PostDevicesBindResponse, error) {
	resp := &PostDevicesBindResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type PostDevicesBindResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		TuyaDeviceId string `json:"tuya_device_id"` // 涂鸦设备id
		TuyaUserId   string `json:"tuya_user_id"`   // 涂鸦用户id
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

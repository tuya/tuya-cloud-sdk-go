package device

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetDeviceFunctionByCategoryReq struct {
	Category string //device category
}

func (t *GetDeviceFunctionByCategoryReq) Method() string {
	return common.RequestGet
}

func (t *GetDeviceFunctionByCategoryReq) API() string {
	return fmt.Sprintf("/v1.0/functions/%s", t.Category)
}

//GetDeviceFunctionByCategory Obtain function list based on category
func GetDeviceFunctionByCategory(category string) (*GetDeviceFunctionResponse, error) {
	a := &GetDeviceFunctionByCategoryReq{Category: category}
	resp := &GetDeviceFunctionResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetDeviceFunctionByCategoryResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		Category  string `json:"category"`
		Functions []struct {
			Name   string `json:"name"`
			Desc   string `json:"desc"`
			Code   string `json:"code"`
			Type   string `json:"type"`
			Values string `json:"values"`
		} `json:"functions"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

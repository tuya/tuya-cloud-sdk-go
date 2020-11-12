package infrared

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetRemotesReq struct {
	InfraredID string
}

func (t *GetRemotesReq) Method() string {
	return common.RequestGet
}

func (t *GetRemotesReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/remotes", t.InfraredID)
}

// GetRemotes sets infrared learning mode
func GetRemotes(infraredID string) (*GetRemotesResponse, error) {
	a := &GetRemotesReq{InfraredID: infraredID}
	resp := &GetRemotesResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetRemotesResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		CategoryID   string `json:"category_id"`
		BrandID      string `json:"brand_id"`
		BrandName    string `json:"brand_name"`
		RemoteID     string `json:"remote_id"`
		RemoteName   string `json:"remote_name"`
		RemoteIndex  string `json:"remote_index"`
		OperatorID   string `json:"operator_id"`
		OperatorName string `json:"operator_name"`
		AreaID       string `json:"area_id"`
		AreaName     string `json:"area_name"`
		IPTVType     int    `json:"iptv_type"`
		T            int64  `json:"t"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

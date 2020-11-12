package infrared

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetBrandsByRemoteIndexReq struct {
	InfraredID  string
	CategoryID  int
	RemoteIndex int
	Region      string
	Lang        string
}

func (t *GetBrandsByRemoteIndexReq) Method() string {
	return common.RequestGet
}

func (t *GetBrandsByRemoteIndexReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/categories/%d/remotes/%d/brands?region=%s", t.InfraredID, t.CategoryID, t.RemoteIndex, t.Region)
}

// GetBrandsByRemoteIndex sets infrared learning mode
func GetBrandsByRemoteIndex(infraredID string, categoryId, remoteIndex int, region, lang string) (*GetBrandsByRemoteIndexResponse, error) {
	a := &GetBrandsByRemoteIndexReq{
		InfraredID:  infraredID,
		CategoryID:  categoryId,
		RemoteIndex: remoteIndex,
		Region:      region,
		Lang:        lang,
	}
	resp := &GetBrandsByRemoteIndexResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetBrandsByRemoteIndexResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  []struct {
		BrandID   int    `json:"brand_id"`
		BrandName string `json:"brand_name"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

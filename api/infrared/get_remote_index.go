package infrared

import (
"fmt"

"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

type GetRemoteIndexesReq struct {
	InfraredID string
	CategoryID string
	BrandID    string
}

func (t *GetRemoteIndexesReq) Method() string {
	return common.RequestGet
}

func (t *GetRemoteIndexesReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/categories/%s/brands/%s", t.InfraredID, t.CategoryID, t.BrandID)
}

// GetRemoteIndexes - get remote control index by supported brands.
func GetRemoteIndexes(infraredID, categoryID, brandID string) (*GetRemoteIndexesResponse, error) {
	a := &GetRemoteIndexesReq{InfraredID: infraredID, CategoryID: categoryID, BrandID: brandID}
	resp := &GetRemoteIndexesResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetRemoteIndexesResponse struct {
	Success bool `json:"success"`
	Result  []struct {
		RemoteIndex   string `json:"remote_index"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
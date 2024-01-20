package infrared

import (
"fmt"

"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

type GetBrandsReq struct {
	InfraredID string
	CategoryID string
}

func (t *GetBrandsReq) Method() string {
	return common.RequestGet
}

func (t *GetBrandsReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/categories/%s/brands", t.InfraredID, t.CategoryID)
}

// GetBrands - get supported brand list.
func GetBrands(infraredID, categoryID string) (*GetBrandsResponse, error) {
	a := &GetBrandsReq{InfraredID: infraredID, CategoryID: categoryID}
	resp := &GetBrandsResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetBrandsResponse struct {
	Success bool `json:"success"`
	Result  []struct {
		BrandID   string `json:"brand_id"`
		BrandName string `json:"brand_name"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
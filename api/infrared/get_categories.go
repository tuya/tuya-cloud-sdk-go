package infrared

import (
	"fmt"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

type GetCategoriesReq struct {
	InfraredID string
}

func (t *GetCategoriesReq) Method() string {
	return common.RequestGet
}

func (t *GetCategoriesReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/categories", t.InfraredID)
}

// GetCategories - get IR device categories
func GetCategories(infraredID string) (*GetCategoriesResponse, error) {
	a := &GetCategoriesReq{InfraredID: infraredID}
	resp := &GetCategoriesResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetCategoriesResponse struct {
	Success bool `json:"success"`
	Result  []struct {
		CategoryID   string `json:"category_id"`
		CategoryName string `json:"category_name"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

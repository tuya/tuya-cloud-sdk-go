package infrared

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetPairingRulesReq struct {
	InfraredID  string
	CategoryID  int
	BrandID     int
	RemoteIndex int
}

func (t *GetPairingRulesReq) Method() string {
	return common.RequestGet
}

func (t *GetPairingRulesReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/categories/%d/brands/%d/remotes%d/rules",
		t.InfraredID, t.CategoryID, t.BrandID, t.RemoteIndex)
}

// GetPairingRules sets infrared learning mode
func GetPairingRules(infraredID string, categoryId, brandID, remoteIndex int) (*GetPairingRulesResponse, error) {
	a := &GetPairingRulesReq{
		InfraredID:  infraredID,
		CategoryID:  categoryId,
		BrandID:     brandID,
		RemoteIndex: remoteIndex,
	}
	resp := &GetPairingRulesResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetPairingRulesResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  []struct {
		KeyName string `json:"key_name"`
		Key     string `json:"key"`
		Desc    string `json:"desc"`
		Code    string `json:"code"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

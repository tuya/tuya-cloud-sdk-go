package infrared

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetMatchingRemotesReq struct {
	InfraredID string
	Token      string
}

func (t *GetMatchingRemotesReq) Method() string {
	return common.RequestGet
}

func (t *GetMatchingRemotesReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/matching-remotes?token=%s", t.InfraredID, t.Token)
}

// GetMatchingRemotes sets infrared learning mode
func GetMatchingRemotes(infraredID, token string) (*GetMatchingRemotesResponse, error) {
	a := &GetMatchingRemotesReq{InfraredID: infraredID, Token: token}
	resp := &GetMatchingRemotesResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetMatchingRemotesResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		HasNext      bool   `json:"has_next"`
		Progress     string `json:"progress"`
		RemoteIndexs []int  `json:"remote_indexs"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

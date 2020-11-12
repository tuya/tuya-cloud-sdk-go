package infrared

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type GetACStatusReq struct {
	InfraredID string
	RemoteID   string
}

func (t *GetACStatusReq) Method() string {
	return common.RequestGet
}

func (t *GetACStatusReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/remotes/%s/ac/status", t.InfraredID, t.RemoteID)
}

// GetACStatus sets infrared learning mode
func GetACStatus(infraredID, remoteID string) (*GetACStatusResponse, error) {
	a := &GetACStatusReq{InfraredID: infraredID, RemoteID: remoteID}
	resp := &GetACStatusResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetACStatusResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		Mode     string `json:"mode"`
		Temp     string `json:"temp"`
		Wind     string `json:"wind"`
		Power    string `json:"power"`
		RemoteID string `json:"remote_id"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

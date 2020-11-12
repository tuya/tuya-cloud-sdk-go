package infrared

import (
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type DeleteRemoteReq struct {
	InfraredID string
	RemoteID   string
}

func (t *DeleteRemoteReq) Method() string {
	return common.RequestDelete
}

func (t *DeleteRemoteReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/remotes/%s", t.InfraredID, t.RemoteID)
}

// DeleteRemote sends test command to infrared controlled device
func DeleteRemote(
	infraredID string,
	remoteID string,
) (*DeleteRemoteResponse, error) {
	a := &DeleteRemoteReq{
		InfraredID: infraredID,
		RemoteID:   remoteID,
	}
	resp := &DeleteRemoteResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type DeleteRemoteResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  bool  `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

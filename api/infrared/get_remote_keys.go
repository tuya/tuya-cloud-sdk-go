package infrared

import (
	"fmt"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

type GetRemoteKeysReq struct {
	InfraredID string
	RemoteID   string
}

func (t *GetRemoteKeysReq) Method() string {
	return common.RequestGet
}

func (t *GetRemoteKeysReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/remotes/%s/keys", t.InfraredID, t.RemoteID)
}

// GetRemoteKeys get keys list by RemoteID
func GetRemoteKeys(infraredID string, remoteID string) (*GetRemoteKeysResponse, error) {
	a := &GetRemoteKeysReq{InfraredID: infraredID, RemoteID: remoteID}
	resp := &GetRemoteKeysResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type GetRemoteKeysResponse struct {
	Success bool `json:"success"`
	Result  struct {
		CategoryID     string `json:"category_id"`
		BrandID        string `json:"brand_id"`
		RemoteIndex    string `json:"remote_index"`
		SingleAir      bool   `json:"single_air"`
		DuplicatePower bool   `json:"duplicate_power"`
		KeyList        []struct {
			Key         string `json:"key"`
			KeyID       int    `json:"key_id"`
			KeyName     string `json:"key_name"`
			StandardKey bool   `json:"standard_key"`
		} `json:"key_list"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

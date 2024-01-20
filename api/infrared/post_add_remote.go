package infrared

import (
	"encoding/json"
	"fmt"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

type PostAddRemoteReq struct {
	InfraredID  string
	CategoryID  string `json:"category_id"`
	BrandID     string `json:"brand_id"`
	BrandName   string `json:"brand_name"`
	RemoteIndex string `json:"remote_index"`
	RemoteName  string `json:"remote_name"`
	// Next parameters for STB supported in China only
	OperatorID   string `json:"operator_id"`
	OperatorName string `json:"operator_name"`
	AreaID       string `json:"area_id"`
	AreaName     string `json:"area_name"`
	IPTVType     int    `json:"iptv_type"`
}

func (t *PostAddRemoteReq) Method() string {
	return common.RequestPost
}

func (t *PostAddRemoteReq) API() string {
	return fmt.Sprintf("/v1.0/infrareds/%s/add-remote", t.InfraredID)
}

func (t *PostAddRemoteReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PostAddRemote sends test command to infrared controlled device
func PostAddRemote(
	infraredID string,
	categoryID string,
	brandID string,
	brandName string,
	remoteIndex string,
	remoteName string,
	operatorID string,
	operatorName string,
	areaID string,
	areaName string,
	iptvType int,
) (*PostAddRemoteResponse, error) {
	a := &PostAddRemoteReq{
		InfraredID:   infraredID,
		CategoryID:   categoryID,
		BrandID:      brandID,
		BrandName:    brandName,
		RemoteIndex:  remoteIndex,
		RemoteName:   remoteName,
		OperatorID:   operatorID,
		OperatorName: operatorName,
		AreaID:       areaID,
		AreaName:     areaName,
		IPTVType:     iptvType,
	}
	resp := &PostAddRemoteResponse{}
	err := common.DoAPIRequest(a, resp)
	return resp, err
}

type PostAddRemoteResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		RemoteID string `json:"remote_id"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

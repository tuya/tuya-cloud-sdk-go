package user

import (
	"encoding/json"
	"fmt"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

type PostUserRegisterReq struct {
	Schema       string // If your app is based on a Tuya-based SDK, you can query the channel identifier under the SDK app details page on the platform.
	CountryCode  string `json:"country_code"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	NickName     string `json:"nick_name"`
	UsernameType string `json:"username_type"` // username type,1:mobile,2:email,3:username, default username
}

func (t *PostUserRegisterReq) Method() string {
	return common.RequestPost
}

func (t *PostUserRegisterReq) API() string {
	return fmt.Sprintf("/v1.0/apps/%s/user", t.Schema)
}

func (t *PostUserRegisterReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PostUserRegister  Register users in App corresponding to schema
func PostUserRegister(req *PostUserRegisterReq) (*PostUserRegisterResponse, error) {
	resp := &PostUserRegisterResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

// PostUserRegisterBy Register users in App corresponding to schema
func PostUserRegisterBy(Schema string,
	CountryCode, Username, Password, NickName, UsernameType string) (*PostUserRegisterResponse, error) {
	req := &PostUserRegisterReq{Schema: Schema,
		CountryCode:  CountryCode,
		Username:     Username,
		Password:     Password,
		NickName:     NickName,
		UsernameType: UsernameType}
	resp := &PostUserRegisterResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type PostUserRegisterResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		UID string `json:"uid"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

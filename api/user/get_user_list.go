package user

import "fmt"
import "github.com/TuyaInc/tuya_cloud_sdk_go/api/common"

type GetUserListReq struct {
	Schema   string // The schema is the channel identifier of the APP application.
	PageNo   uint
	PageSize uint
}

func (t *GetUserListReq) Method() string {
	return common.RequestGet
}

func (t *GetUserListReq) API() string {
	return fmt.Sprintf(
		"/v1.0/apps/%s/users?page_no=%d&page_size=%d",
		t.Schema,
		t.PageNo,
		t.PageSize)
}

// GetUserList Obtain App users based on schema page
func GetUserList(schema string, pageNo, pageSize uint) (*GetUserListResponse, error) {
	req := &GetUserListReq{Schema: schema,
		PageNo:   pageNo,
		PageSize: pageSize}
	resp := &GetUserListResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type GetUserListResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		List []struct {
			UID      string `json:"uid"`
			Username string `json:"username"`
		} `json:"list"`
		HasMore bool `json:"has_more"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

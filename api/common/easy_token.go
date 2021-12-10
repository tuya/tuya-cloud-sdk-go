package common

import (
	"github.com/tuya/tuya-cloud-sdk-go/pkg/tylog"
)

type GetTokenReq struct {
}

func (t *GetTokenReq) Method() string {
	return RequestGet
}

func (t *GetTokenReq) API() string {
	return "/v1.0/token?grant_type=1"
}

func GetTokenAPI() (*GetTokenAPIResponse, error) {
	getTokenReq := &GetTokenReq{}
	req, err := NewHTTPRequest(getTokenReq)
	if err != nil {
		tylog.SugarLog.Infof("GetTokenAPI failed err:%v,req:%v\n", err, req)
		return nil, err
	}

	timestamp := GetTimestamp()
	AddEasyHeader(req, timestamp)

	resp := &GetTokenAPIResponse{}
	err = DoRequest(req, resp)
	if err != nil {
		tylog.SugarLog.Infof("GetTokenAPI failed err:%v,req:%v,resp:%v\n", err, req, resp)
		return nil, err
	}
	SetToken(resp.Result.AccessToken, resp.Result.RefreshToken, resp.Result.ExpireTime)
	return resp, nil
}

type GetTokenAPIResponse struct {
	Success bool `json:"success"`
	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`

	T      int64 `json:"t"`
	Result struct {
		ExpireTime   int    `json:"expire_time"`
		UID          string `json:"uid"`
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	} `json:"result"`
}

package common

import (
	"fmt"

	"github.com/tuya/tuya-cloud-sdk-go/pkg/tylog"
)

type RefreshTokenReq struct {
	curRefreshToken string
}

func (t *RefreshTokenReq) Method() string {
	return RequestGet
}

func (t *RefreshTokenReq) API() string {
	return fmt.Sprintf("/v1.0/token/%s", t.curRefreshToken)
}

// DoRefreshToken  To refresh token
func DoRefreshToken() (*RefreshTokenResponse, error) {
	refreshToken, err := GetRefreshToken()
	if err != nil || len(refreshToken) == 0 {
		tylog.SugarLog.Infof("GetRefreshToken failed err:%v\n", err)
		return nil, err
	}
	refreshTokenReq := &RefreshTokenReq{}
	refreshTokenReq.curRefreshToken = refreshToken

	req, err := NewHTTPRequest(refreshTokenReq)
	if err != nil {
		tylog.SugarLog.Infof("DoRefreshToken failed err:%v,req:%v\n", err, req)
		return nil, err
	}
	timestamp := GetTimestamp()
	AddEasyHeader(req, timestamp)

	resp := &RefreshTokenResponse{}
	err = DoRequest(req, resp)
	if err != nil {
		tylog.SugarLog.Infof("DoRefreshToken failed err:%v,req:%v,resp:%v\n", err, req, resp)
		return nil, err
	}

	SetToken(resp.Result.AccessToken, resp.Result.RefreshToken, resp.Result.ExpireTime)
	return resp, nil
}

type RefreshTokenResponse struct {
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

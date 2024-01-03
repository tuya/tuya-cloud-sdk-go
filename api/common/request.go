package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/tuya/tuya-cloud-sdk-go/config"

	"github.com/tuya/tuya-cloud-sdk-go/pkg/tylog"
)

func DoAPIRequest(a APIRequest, resp interface{}) error {
	var token, err = TokenLocalCache.GetToken()
	if err != nil {
		return ErrorGetTokenFailed
	}

	uri := strings.Join([]string{config.HOST, a.API()}, "")
	timestamp := GetTimestamp()
	var req *http.Request
	pr, ok := a.(RequestBody)
	if ok {
		req, err = http.NewRequest(a.Method(), uri, bytes.NewReader(pr.Body()))
	} else {
		req, err = http.NewRequest(a.Method(), uri, nil)
	}
	if err != nil {
		return err
	}
	if a.Method() != "GET" {
		AddBodyBizHeader(req, token, timestamp)
	} else {
		AddBizHeader(req, token, timestamp)
	}

	err = DoRequest(req, resp)
	return err
}

func AddEasyHeader(req *http.Request, timestamp string) {
	req.Header.Add("client_id", config.AccessID)
	req.Header.Add("sign_method", "HMAC-SHA256")
	req.Header.Add("t", timestamp)
	sign := GetEasySignV2(req)
	req.Header.Set("sign", sign)
}

func AddBizHeader(req *http.Request, token, timestamp string) {
	req.Header.Add("client_id", config.AccessID)
	req.Header.Add("access_token", token)
	req.Header.Add("sign_method", "HMAC-SHA256")
	req.Header.Add("t", timestamp)
	sign := GetBizSignV2(req, token)
	req.Header.Set("sign", sign)
}

func AddBodyBizHeader(req *http.Request, token, timestamp string) {
	req.Header.Add("client_id", config.AccessID)
	req.Header.Add("access_token", token)
	req.Header.Add("sign_method", "HMAC-SHA256")
	req.Header.Add("t", timestamp)
	req.Header.Add("Content-Type", "application/json")
	sign := GetBizSignV2(req, token)
	req.Header.Set("sign", sign)
}

func DoRequest(req *http.Request, resp interface{}) error {
	httpResp, err := http.DefaultClient.Do(req)
	if err != nil {
		tylog.SugarLog.Errorf("do request failed err:%v,req:%v\n", err, req)
		return err
	}
	defer httpResp.Body.Close()
	bs, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		tylog.SugarLog.Errorf("do request failed err:%v,req:%v\n", err, req)
		return err
	}

	// resp := &GetFunctionsResponse{}
	err = json.Unmarshal(bs, &resp)
	if err != nil {
		tylog.SugarLog.Errorf("do request failed err:%v,req:%v,resp:%v\n", err, req, string(bs))
		return err
	}
	handlerError(bs)
	tylog.SugarLog.Infof("req:%v,resp:%+v\n", req, resp)
	return nil
}

type ErrorInfo struct {
	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

const TokenInvalid int = 1010

func handlerError(bs []byte) {
	e := ErrorInfo{}
	json.Unmarshal(bs, &e)
	if e.Code == TokenInvalid {
		GetTokenAPI()
	}
}

func NewHTTPRequest(a APIRequest) (*http.Request, error) {
	url := strings.Join([]string{config.HOST, a.API()}, "")
	var req *http.Request
	var err error
	pr, ok := a.(RequestBody)
	if ok {
		req, err = http.NewRequest(a.Method(), url, bytes.NewReader(pr.Body()))
	} else {
		req, err = http.NewRequest(a.Method(), url, nil)
	}
	return req, err
}

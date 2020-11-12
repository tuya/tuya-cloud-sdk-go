package device_dn

import "encoding/json"
import "github.com/TuyaInc/tuya_cloud_sdk_go/api/common"

type PostDeviceTokenReq struct {
	UID        string `json:"uid"`            // 用户唯一标识	是
	TimeZoneID string `json:"timeZoneId"`     // 用户所在时区id，州/省份（Asia/Shanghai）	是
	Lon        string `json:"lon,omitempty"`  // 经度	否
	Lat        string `json:"lat,omitempty"`  // 纬度	否
	Lang       string `json:"lang,omitempty"` // 系统语言，zh,eu等，默认zh	否
}

func (t *PostDeviceTokenReq) Method() string {
	return common.RequestPost
}

func (t *PostDeviceTokenReq) API() string {
	return "/v1.0/devices/token"
}

func (t *PostDeviceTokenReq) Body() []byte {
	bs, _ := json.Marshal(t)
	return bs
}

// PostDeviceTokenBy Generate device distribution network token
func PostDeviceTokenBy(uid, timeZoneID, lon, lat, lang string) (*PostDeviceTokenResponse, error) {
	req := &PostDeviceTokenReq{
		UID:        uid,
		TimeZoneID: timeZoneID,
		Lon:        lon,
		Lat:        lat,
		Lang:       lang,
	}
	resp := &PostDeviceTokenResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

// PostDeviceToken Generate device distribution network token
func PostDeviceToken(req *PostDeviceTokenReq) (*PostDeviceTokenResponse, error) {
	resp := &PostDeviceTokenResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type PostDeviceTokenResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		Secret string `json:"secret"`
		Region string `json:"region"`
		Token  string `json:"token"`
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

package pairing

import (
	"encoding/json"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

type ExtensionReq struct {
	UUID string `json:"uuid"` // Universally unique ID of the device
}

type PostPairingTokenReq struct {
	PairingType string       `json:"pairing_type"`        // Pairing type, including BLE, AP, and EZ
	UID         string       `json:"uid"`                 // 	Unique user identifier
	TimeZoneId  string       `json:"time_zone_id"`        // The ID of the user’s time zone, required for daylight saving time
	HomeID      string       `json:"home_id,omitempty"`   // Family ID. If left blank, it refers to the user’s default family
	Extension   ExtensionReq `json:"extension,omitempty"` // Extension information. When the pairing type is BLE, the device UUID must be passed in. This parameter can be ignored in case of Wi-Fi pairing
}

func (t *PostPairingTokenReq) Method() string {
	return common.RequestPost
}

func (t *PostPairingTokenReq) API() string {
	return "/v1.0/device/paring/token"
}

func (t *PostPairingTokenReq) Body() []byte {
	reqBody, _ := json.Marshal(t)
	return reqBody
}

// PostPairingToken Obtain pairing information
func PostPairingToken(PairingType, UID, TimeZoneID, HomeID, ExtensionUUID string) (*PostPairingTokenResponse, error) {
	req := &PostPairingTokenReq{
		PairingType: PairingType,
		UID:         UID,
		TimeZoneId:  TimeZoneID,
		HomeID:      HomeID,
		Extension:   ExtensionReq{},
	}
	if ExtensionUUID != "" {
		req.Extension.UUID = ExtensionUUID
	}
	resp := &PostPairingTokenResponse{}
	err := common.DoAPIRequest(req, resp)
	return resp, err
}

type ExtensionResp struct {
	EncryptKey string `json:"encrypt_key"` // Encryption key
	Random     string `json:"random"`      // Encrypted string
}

type PostPairingTokenResponse struct {
	Success bool  `json:"success"`
	T       int64 `json:"t"`
	Result  struct {
		ExpireTime int           `json:"expire_time"` // Expiration time of the token
		Region     string        `json:"region"`      // Current available zone, including `AY`, `EU`, and `US`
		Token      string        `json:"token"`       // Pairing token
		Secret     string        `json:"secret"`      // Key
		Extension  ExtensionResp `json:"extension"`   // Extension information
	} `json:"result"`

	// error info
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

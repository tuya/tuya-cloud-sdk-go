package common

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/TuyaInc/tuya_cloud_sdk_go/config"
)

func GetBizSign(token, timestamp string) string {
	sign := strings.ToUpper(HmacSha256(config.AccessID+token+timestamp, config.AccessKey))
	return sign
}

func GetEasySign(timestamp string) string {
	sign := strings.ToUpper(HmacSha256(config.AccessID+timestamp, config.AccessKey))
	return sign
}

func HmacSha256(data, key string) string {

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(key))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))
	return sha
}

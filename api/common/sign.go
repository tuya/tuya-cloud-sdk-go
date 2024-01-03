package common

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"

	"github.com/tuya/tuya-cloud-sdk-go/config"
)

func GetBizSign(token, timestamp string) string {
	sign := strings.ToUpper(HmacSha256(config.AccessID+token+timestamp, config.AccessKey))
	return sign
}

func GetEasySign(timestamp string) string {
	sign := strings.ToUpper(HmacSha256(config.AccessID+timestamp, config.AccessKey))
	return sign
}

func GetBizSignV2(req *http.Request, token string) string {
	contentSha256 := ""
	if req.Body != nil {
		buf, _ := ioutil.ReadAll(req.Body)
		req.Body = ioutil.NopCloser(bytes.NewBuffer(buf))
		contentSha256 = GetSha256(buf)
	} else {
		contentSha256 = GetSha256([]byte(""))
	}

	headers := ""
	signHeaderKeys := req.Header.Get("Signature-Headers")
	if signHeaderKeys != "" {
		keys := strings.Split(signHeaderKeys, ":")
		for _, key := range keys {
			headers += key + ":" + req.Header.Get(key) + "\n"
		}
	}

	uri := req.URL.Path
	keys := make([]string, 0, 10)
	form, err := url.ParseQuery(req.URL.RawQuery)
	if err == nil {
		for key, _ := range form {
			keys = append(keys, key)
		}
	}
	if len(keys) > 0 {
		uri += "?"
		sort.Strings(keys)
		for _, keyName := range keys {
			value := form.Get(keyName)
			uri += keyName + "=" + value + "&"
		}
		uri = strings.TrimSuffix(uri, "&")
	}

	stringToSign := req.Method + "\n" + contentSha256 + "\n" + headers + "\n" + uri
	nonce := req.Header.Get("nonce")
	t := req.Header.Get("t")
	str := config.AccessID + token + t + nonce + stringToSign
	sign := strings.ToUpper(HmacSha256(str, config.AccessKey))
	return sign
}

func GetEasySignV2(req *http.Request) string {
	sign := GetBizSignV2(req, "")
	return sign
}

func GetSha256(data []byte) string {
	sha256Contain := sha256.New()
	sha256Contain.Write(data)
	return hex.EncodeToString(sha256Contain.Sum(nil))
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

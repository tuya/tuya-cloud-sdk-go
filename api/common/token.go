package common

import (
	"sync"
	"time"

	"github.com/TuyaInc/tuya_cloud_sdk_go/pkg/tylog"
)

type Info struct {
	Token        string
	RefreshToken string
	ExpireAt     time.Time
	Mu           *sync.RWMutex
}

var LocalToken = Info{Mu: &sync.RWMutex{}}

func SetToken(token, refreshToken string, expire int) {
	LocalToken.Mu.Lock()
	LocalToken.Token = token
	LocalToken.RefreshToken = refreshToken
	LocalToken.ExpireAt = time.Now().Add(time.Duration(expire) * time.Second)
	LocalToken.Mu.Unlock()
}

func GetRefreshToken() (string, error) {
	if LocalToken.RefreshToken != "" {
		return LocalToken.RefreshToken, nil
	}
	_, err := GetToken()
	if err != nil {
		return "", err
	}
	return LocalToken.RefreshToken, nil
}

func GetToken() (string, error) {
	LocalToken.Mu.RLock()
	t := LocalToken.Token
	expire := LocalToken.ExpireAt
	LocalToken.Mu.RUnlock()
	// token不为空，且有效时间大于30秒，直接返回缓存token
	if t != "" && expire.After(time.Now().Add(30*time.Second)) {
		return t, nil
	}
	tylog.SugarLog.Info("without token, the token will be pulled again")
	if t == "" || LocalToken.RefreshToken == "" {
		_, err := GetTokenAPI()
		if err != nil {
			return "", err
		}
	} else { // 小于30s就刷新下
		// token不为空，调用RefreshTokenAPI刷新token
		_, err := DoRefreshToken()
		if err != nil {
			return "", err
		}
	}

	LocalToken.Mu.RLock()
	t = LocalToken.Token
	LocalToken.Mu.RUnlock()
	return t, nil
}

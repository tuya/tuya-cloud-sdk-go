package common

import (
	"github.com/tuya/tuya-cloud-sdk-go/pkg/tylog"
	"sync"
	"time"
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
	// token already exists
	if t != "" {
		if expire.Sub(time.Now())<=0{
			// token is expired
			_, err := GetTokenAPI()
			if err != nil {
				return "", err
			}
		}else if expire.Sub(time.Now())>0 && expire.Sub(time.Now())<=30*time.Second {
			// token will expire after 30s
			_, err := DoRefreshToken()
			if err != nil {
				return "", err
			}
		}else{
			return t, nil
		}
	}

	tylog.SugarLog.Info("without token, the token will be pulled again")
	_, err := GetTokenAPI()
	if err != nil {
		return "", err
	}

	LocalToken.Mu.RLock()
	t = LocalToken.Token
	LocalToken.Mu.RUnlock()
	return t, nil
}

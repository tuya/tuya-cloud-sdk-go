package common

type APIRequest interface {
	Method() string
	API() string
}

type RequestBody interface {
	Body() []byte
}

// TokenLocalManage token of local cache manage
type TokenManage interface {
	GetToken() (string, error)
}

type tokenLocalCache struct {
}

func (t *tokenLocalCache) GetToken() (string, error) {
	return GetToken()
}

var TokenLocalCache TokenManage = &tokenLocalCache{}

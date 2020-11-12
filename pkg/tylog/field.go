package tylog

import "go.uber.org/zap"

func String(key string, value string) zap.Field {
	return zap.String(key, value)
}

func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

func ErrorField(err error) zap.Field {
	return zap.Error(err)
}

func InfoField(v interface{}) zap.Field {
	return zap.Any("info", v)
}

// ByteString 会把[]byte转成string
// 注意 Any() 会把[]byte转成 binary
func ByteString(key string, val []byte) zap.Field {
	return zap.ByteString(key, val)
}

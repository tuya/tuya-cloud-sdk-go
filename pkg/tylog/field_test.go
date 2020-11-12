package tylog

import (
	"errors"
	"testing"

	"go.uber.org/zap"
)

func TestString(t *testing.T) {
	key := "test_key"
	value := "test_value"
	f := String(key, value)
	SetGlobalLog("test", true)
	allLevelWithField(f)
}

func TestAny(t *testing.T) {
	tables := []struct {
		key   string
		value interface{}
	}{
		{"test_key", "test_value_str"},
		{"test_key", []byte("test_value_byte")},
		{"test_key", 666},
	}

	SetGlobalLog("test", true)
	for _, v := range tables {
		f := Any(v.key, v.value)
		allLevelWithField(f)
	}
}

func TestErrorField(t *testing.T) {
	value := "test_value"
	e := errors.New(value)
	f := ErrorField(e)
	SetGlobalLog("test", true)
	allLevelWithField(f)
}

func TestInfoField(t *testing.T) {
	value := "test_value"
	f := InfoField(value)
	SetGlobalLog("test", true)
	allLevelWithField(f)
}

func TestByteString(t *testing.T) {
	key := "test_key"
	value := []byte("test_value")
	f := ByteString(key, value)
	SetGlobalLog("test", true)
	allLevelWithField(f)
}

func allLevelWithField(f zap.Field) {
	Debug("debug", f)
	Info("info", f)
	Warn("warn", f)
	Error("error", f)
}

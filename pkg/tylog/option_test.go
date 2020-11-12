package tylog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithLocalTimeOption(t *testing.T) {
	cfg := &LogConfig{}
	WithLocalTimeOption(true)(cfg)
	assert.Equal(t, true, cfg.localTime)
}

func TestWithMultiWriteOption(t *testing.T) {
	cfg := &LogConfig{}
	v := true
	WithMultiWriteOption(v)(cfg)
	assert.Equal(t, v, cfg.multiWrite)
}

func TestWithMaxSizeOption(t *testing.T) {
	cfg := &LogConfig{}
	v := 10
	WithMaxSizeOption(v)(cfg)
	assert.Equal(t, v, cfg.maxSize)
}

func TestWithRotatePeriodSecondOption(t *testing.T) {
	cfg := &LogConfig{}
	v := int64(10)
	WithRotatePeriodSecondOption(v)(cfg)
	assert.Equal(t, v, cfg.rotatePeriodSecond)
}

func TestWithLevelOption(t *testing.T) {
	cfg := &LogConfig{}
	v := LevelEnumInfo
	WithLevelOption(v)(cfg)
	assert.Equal(t, v, cfg.level)
}

func TestWithAppNameOption(t *testing.T) {
	cfg := &LogConfig{}
	v := "test"
	WithAppNameOption(v)(cfg)
	assert.Equal(t, v, cfg.appName)
}

func TestWithDirOption(t *testing.T) {
	cfg := &LogConfig{}
	v := "logDir"
	WithDirOption(v)(cfg)
	assert.Equal(t, v, cfg.dir)
}

func TestWithFormatOption(t *testing.T) {
	cfg := &LogConfig{}
	v := FormatEnumJSON
	WithFormatOption(v)(cfg)
	assert.Equal(t, v, cfg.format)
}

func TestWithHooksOption(t *testing.T) {
	cfg := &LogConfig{}
	v := []LogHook{hookImpl{}}
	WithHooksOption(v...)(cfg)
	assert.Equal(t, v, cfg.Hooks)
}

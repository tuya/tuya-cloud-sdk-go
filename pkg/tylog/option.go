package tylog

import "go.uber.org/zap/zapcore"

type OptionFunc func(*LogConfig)

type LogHook interface {
	DoHook(zapcore.Entry) error
}

func WithLocalTimeOption(v bool) OptionFunc {
	return OptionFunc(func(log *LogConfig) {
		log.localTime = v
	})
}

func WithMultiWriteOption(v bool) OptionFunc {
	return OptionFunc(func(log *LogConfig) {
		log.multiWrite = v
	})
}

func WithMaxSizeOption(v int) OptionFunc {
	return OptionFunc(func(log *LogConfig) {
		log.maxSize = v
	})
}

func WithRotatePeriodSecondOption(v int64) OptionFunc {
	return OptionFunc(func(log *LogConfig) {
		log.rotatePeriodSecond = v
	})
}

func WithLevelOption(v LevelEnum) OptionFunc {
	return OptionFunc(func(log *LogConfig) {
		log.level = v
	})
}

func WithAppNameOption(v string) OptionFunc {
	return OptionFunc(func(log *LogConfig) {
		log.appName = v
	})
}

func WithDirOption(v string) OptionFunc {
	return OptionFunc(func(log *LogConfig) {
		log.dir = v
	})
}

func WithFormatOption(v FormatEnum) OptionFunc {
	return OptionFunc(func(log *LogConfig) {
		log.format = v
	})
}

func WithHooksOption(hooks ...LogHook) OptionFunc {
	return OptionFunc(func(log *LogConfig) {
		log.Hooks = hooks
	})
}

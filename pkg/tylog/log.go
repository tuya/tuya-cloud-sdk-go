package tylog

import (
	"time"

	"github.com/TuyaInc/tuya_cloud_sdk_go/pkg/tyutils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var log *zap.Logger

var SugarLog *zap.SugaredLogger

func SetGlobalLog(appName string, prod bool, opts ...OptionFunc) {
	if appName == "" {
		panic("appName is not allowed to be empty")
	}
	var config *LogConfig
	if prod {
		config = ProdLogConfig(appName)
	} else {
		config = DefaultLogConfig()
		config.appName = appName
	}

	for _, v := range opts {
		v(config)
	}

	log = NewLog(config)
	SugarLog = log.Sugar()
}

func NewLog(config *LogConfig) *zap.Logger {
	if !tyutils.IsDir(config.dir) {
		err := tyutils.Mkdir(config.dir)
		if err != nil {
			panic("mkdir logs failed")
		}
	}
	filePath := config.dir + "/" + config.appName + ".log"

	// lumberjack 实现了一些日志分割的功能
	jLoger := &lumberjack.Logger{
		Filename: filePath,
		MaxSize:  config.maxSize,
	}
	jLoger.LocalTime = config.localTime

	// writer
	var writer zapcore.WriteSyncer
	if config.multiWrite {
		w1 := zapcore.AddSync(jLoger)
		// 日志同时在终端输出
		w2, closeOut, err := zap.Open([]string{"stderr"}...)
		if err != nil {
			if closeOut != nil {
				closeOut()
			}
			panic(err)
		}
		writer = zapcore.NewMultiWriteSyncer(w1, w2)
	} else {
		writer = zapcore.AddSync(jLoger)
	}

	// encoder
	encConfig := zap.NewProductionEncoderConfig()
	encConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	var enc zapcore.Encoder
	if config.format == FormatEnumConsole {
		enc = zapcore.NewConsoleEncoder(encConfig)
	} else {
		enc = zapcore.NewJSONEncoder(encConfig)
	}

	// core
	core := zapcore.NewCore(
		enc,
		writer,
		config.level,
	)

	l := zap.New(core)
	l = l.WithOptions(zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel))

	// hook
	if len(config.Hooks) > 0 {
		hks := make([]zap.Option, 0, len(config.Hooks))
		for _, v := range config.Hooks {
			hk := zap.Hooks(v.DoHook)
			hks = append(hks, hk)

		}
		l = l.WithOptions(hks...)
	}

	RotatePeriod(jLoger, config.rotatePeriodSecond)
	return l
}

func DefaultLogConfig() *LogConfig {
	return &LogConfig{
		localTime:          true,
		multiWrite:         true,
		maxSize:            100,
		rotatePeriodSecond: DefaultRotatePeriodSecond,
		level:              LevelEnumDebug,
		appName:            "tmp",
		dir:                "logs",
		format:             FormatEnumConsole,
	}
}

func ProdLogConfig(appName string) *LogConfig {
	c := DefaultLogConfig()
	c.multiWrite = false
	c.level = LevelEnumInfo
	c.appName = appName
	c.format = FormatEnumJSON
	return c
}

// 日志滚动周期
func RotatePeriod(l *lumberjack.Logger, periodSecond int64) {
	now := time.Now().Unix()
	tickSecond := periodSecond - now%periodSecond
	tk := time.NewTicker(time.Duration(tickSecond) * time.Second)
	go func() {
		<-tk.C
		err := l.Rotate()
		if err != nil {
			Error("rotate failed", ErrorField(err))
		}
		tk.Stop()
		periodSecondTicker := time.NewTicker(time.Duration(periodSecond) * time.Second)
		for range periodSecondTicker.C {
			err := l.Rotate()
			if err != nil {
				Error("rotate failed", ErrorField(err))
			}
		}
	}()
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Debug(msg string, fields ...zap.Field) {
	log.Debug(msg, fields...)
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Info(msg string, fields ...zap.Field) {
	log.Info(msg, fields...)
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Warn(msg string, fields ...zap.Field) {
	log.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func Error(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func Fatal(msg string, fields ...zap.Field) {
	log.Error(msg, fields...)
}

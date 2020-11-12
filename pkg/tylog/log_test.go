package tylog

import (
	"fmt"
	"io/ioutil"
	stdlog "log"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestSetGlobalLog(t *testing.T) {
	SetGlobalLog("test", false)
	Info("info", String("version", "v1"))
}

func TestLogNoSet(t *testing.T) {
	log = new(zap.Logger)
	defer func() {
		err := recover()
		assert.NotNil(t, err)
	}()

	allLevel()
}

func TestSugar(t *testing.T) {
	SetGlobalLog("test", true)
	sugarAllLevel()
}

type hookImpl struct{}

func (hookImpl) DoHook(e zapcore.Entry) error {
	stdlog.Printf("hook e:%v\n", e)
	return nil
}

func TestHook(t *testing.T) {
	SetGlobalLog("test", true, WithHooksOption(hookImpl{}))
	allLevel()
}

func TestRotateByTime(t *testing.T) {
	cmdRemoveDir("logs")
	second := int64(5)
	SetGlobalLog("test", true, WithRotatePeriodSecondOption(second))
	testSecond := int64(20)
	tk := time.NewTicker(time.Duration(testSecond) * time.Second)
	for {
		select {
		case <-tk.C:
			goto end
		default:
			allLevel()
			time.Sleep(100 * time.Millisecond)
		}
	}

end:
	list := cmdLs("logs")
	get, want := int64(len(list)), testSecond/second+1
	assert.Equal(t, want, get)
}

func cmdLs(dir string) []os.FileInfo {
	targetDirPath := dir
	dirList, err := ioutil.ReadDir(targetDirPath)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		for _, dirInfo := range dirList {
			fmt.Println(dirInfo.Name())
		}
	}
	return dirList
}

func cmdRemoveDir(dir string) {
	_ = os.RemoveAll(dir)
}

func TestAtomicLevel(t *testing.T) {
	cmdRemoveDir("logs")
	level := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	SetGlobalLog("test", false, WithLevelOption(level))
	http.Handle("/loglevel", level)
	go func() {
		srv := &http.Server{Addr: ":8080"}
		err := srv.ListenAndServe()
		assert.Nil(t, err)
	}()

	Debug("info", String("version", "v1"))
	req, _ := http.NewRequest("PUT", "http://127.0.0.1:8080/loglevel", strings.NewReader(`{"level":"INFO"}`))
	resp, err := http.DefaultClient.Do(req)
	assert.Nil(t, err)
	bs, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	t.Log("resp:", string(bs))
	defer resp.Body.Close()
	Debug("info", String("version", "v2"))
	// 日志文件中应该只存在第一次Debug的日志
	// 第二次Debug日志因为日志级别已经被调整为info，所以被忽略了
}

func BenchmarkString(b *testing.B) {
	SetGlobalLog("test", true)
	for i := 0; i < b.N; i++ {
		Info("msg", String("key", "value"))
	}
}

func BenchmarkAny(b *testing.B) {
	SetGlobalLog("test", true)
	for i := 0; i < b.N; i++ {
		Info("msg", Any("key", "value"))
	}
}

func BenchmarkSugar(b *testing.B) {
	SetGlobalLog("test", true)
	for i := 0; i < b.N; i++ {
		SugarLog.Infof("msg key:%s", "value")
	}
}

func allLevel() {
	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
}

func sugarAllLevel() {
	SugarLog.Debug("debug", "arg")
	SugarLog.Info("info", "arg")
	SugarLog.Warn("warn", "arg")
	SugarLog.Error("error", "arg")

	SugarLog.Debugf("hi --- %s %d", "debug", 666)
	SugarLog.Infof("hi --- %s %d", "info", 666)
	SugarLog.Warnf("hi --- %s %d", "warn", 666)
	SugarLog.Errorf("hi --- %s %d", "error", 666)
}

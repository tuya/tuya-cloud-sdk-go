package config

import (
	"github.com/TuyaInc/tuya_cloud_sdk_go/pkg/tylog"
)

var HOST string
var AccessID string
var AccessKey string

// SetEnv  you must be init here
func SetEnv(serverHost, accessID, accessKey string) {
	// set log
	SetLog("tysdk", true)
	tylog.SugarLog.Info("init log success")

	// init env
	HOST = serverHost
	AccessID = accessID
	AccessKey = accessKey
}

func SetLog(appName string, prod bool) {
	tylog.SetGlobalLog(appName, prod)
}

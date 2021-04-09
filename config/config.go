package config

import (
	"github.com/tuya/tuya-cloud-sdk-go/pkg/tylog"
)

var HOST string
var AccessID string
var AccessKey string

// SetEnv  you must be init here
func SetEnv(serverHost, accessID, accessKey string) {
	// set log
	//SetLog("tysdk", true)
	//log.Print("init log success")

	// init env
	HOST = serverHost
	AccessID = accessID
	AccessKey = accessKey
}

func SetLog(appName string, prod bool) {
	tylog.SetGlobalLog(appName, prod)
}

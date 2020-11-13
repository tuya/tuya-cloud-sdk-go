package common

import (
	"log"

	"github.com/tuya/tuya_cloud_sdk_go/config"
)

type ExampleData struct {
	TestDataAccessID  string
	TestDataAccessKey string

	TestDataUID        string
	TestDataDeviceID   string
	TestDataTimeZoneID string

	TestDataThirdCloudDeviceId string

	TestDataCountryCode    string
	TestDataAppSchema      string
	TestDataTuyaUsername   string
	TestDataTuyaProductId  string
	TestDataParentDeviceId string

	TestDataPassword string

	TestDataSchema string
}

// test data example
var Ed = ExampleData{}

func SetTestEnv() {
	log.Println("start init")

	config.SetEnv(URLCN, Ed.TestDataAccessID, Ed.TestDataAccessKey)
	if config.HOST == "" || config.AccessID == "" || config.AccessKey == "" {
		log.Println("please set host/accessID/accessKey before run test")
		log.Printf("SetTestEnv failed")
		return
	}
	log.Printf("test env ###### init success,Ed:%v\n", Ed)
}

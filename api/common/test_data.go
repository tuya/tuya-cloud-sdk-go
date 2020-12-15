package common

import (
	"log"

	"github.com/tuya/tuya-cloud-sdk-go/config"
)

type ExampleData struct {
	TestDataAccessID  string
	TestDataAccessKey string

	TestDataUID        string
	TestDataDeviceID   string
	TestDataTimeZoneID string

	TestDataThirdCloudDeviceID string

	TestDataCountryCode    string
	TestDataAppSchema      string
	TestDataTuyaUsername   string
	TestDataTuyaProductID  string
	TestDataParentDeviceID string

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

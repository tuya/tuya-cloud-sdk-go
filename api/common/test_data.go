package common

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/TuyaInc/tuya_cloud_sdk_go/config"
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
	err := loadTestData()
	if err != nil {
		log.Println("load test data failed", err)
	} else {
		log.Println("load test data success")
	}
	config.SetEnv(URLCNPre, Ed.TestDataAccessID, Ed.TestDataAccessKey)
	log.Printf("test env ###### init success,Ed:%v\n", Ed)
}

func loadTestData() error {
	bs, err := ioutil.ReadFile("../../config/test_data.conf")
	if err != nil {
		return err
	}
	if len(bs) > 0 {
		err := json.Unmarshal(bs, &Ed)
		if err != nil {
			return err
		}
	}
	log.Println("test data:", string(bs))
	return nil
}

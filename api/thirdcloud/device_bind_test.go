package thirdcloud

import (
	"fmt"
	"testing"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

func init() {
	common.SetTestEnv()
}

var (
	ThirdCloudDeviceId = common.Ed.TestDataThirdCloudDeviceID
	CountryCode        = common.Ed.TestDataCountryCode
	AppSchema          = common.Ed.TestDataAppSchema
	TuyaUsername       = common.Ed.TestDataTuyaUsername
	TuyaProductId      = common.Ed.TestDataTuyaProductID
	ParentDeviceId     = common.Ed.TestDataParentDeviceID
)

func TestPostDevicesBind(t *testing.T) {
	resp, err := PostDevicesBind(common.Ed.TestDataThirdCloudDeviceID, common.Ed.TestDataCountryCode,
		common.Ed.TestDataAppSchema, common.Ed.TestDataTuyaUsername,
		common.Ed.TestDataTuyaProductID, common.Ed.TestDataParentDeviceID,
	)
	if err != nil {
		t.Errorf("PostDevicesBind req has err:%v,resp:%v \n", err, resp)
	}
	fmt.Printf("resp:%v \n", resp)
	if !resp.Success {
		t.Errorf("resp:%v is not success \n", resp)
	}
}

package thirdcloud

import (
	"fmt"
	"testing"
)

func TestPostDevicesStatus(t *testing.T) {
	status := []PostDevicesStatusItem{
		{
			Code:  "switch_3",
			Value: true,
		},
	}
	resp, err := PostDevicesStatus(ThirdCloudDeviceId, status)
	if err != nil {
		t.Errorf("PostDevicesBind req has err:%v,resp:%v \n", err, resp)
	}
	fmt.Printf("resp:%v \n", resp)
	if !resp.Success {
		t.Errorf("resp:%v is not success \n", resp)
	}
}

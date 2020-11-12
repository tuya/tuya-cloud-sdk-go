package thirdcloud

import (
	"fmt"
	"testing"
)

func TestDeleteDevicesUnBind(t *testing.T) {
	resp, err := DeleteDevicesUnBind(ThirdCloudDeviceId)
	if err != nil {
		t.Errorf("PostDevicesBind req has err:%v,resp:%v \n", err, resp)
	}
	fmt.Printf("resp:%v \n", resp)
	if !resp.Success {
		t.Errorf("resp:%v is not success \n", resp)
	}
}

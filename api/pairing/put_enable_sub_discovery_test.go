package pairing

import (
	"testing"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

func TestEnableSubDiscovery(t *testing.T) {
	reqInfo := &EnableSubDiscoveryReq{
		DeviceId: common.Ed.TestDataDeviceID,
		Duration: 10,
	}
	tests := []struct {
		name    string
		reqInfo *EnableSubDiscoveryReq
		want    *EnableSubDiscoveryResponse
		wantErr bool
	}{
		{
			name:    "1",
			reqInfo: reqInfo,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EnableSubDiscovery(
				tt.reqInfo.DeviceId,
				tt.reqInfo.Duration,
			)
			if (err != nil) != tt.wantErr {
				t.Errorf("SignUp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}

package device

import (
	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"

	"testing"
)

func TestGetDeviceSpecificationsFunction(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		deviceID string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetDeviceSpecResponse
		wantErr bool
	}{

		{
			name: "1",
			args: args{deviceID: deviceID},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDeviceSpecifications(tt.args.deviceID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDeviceSpecifications() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}

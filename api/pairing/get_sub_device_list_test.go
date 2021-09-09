package pairing

import (
	"testing"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

func TestSubDeviceList(t *testing.T) {
	type args struct {
		deviceId string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetPairingDevicesResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{common.Ed.TestDataParentDeviceID},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSubDeviceList(tt.args.deviceId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSubDeviceList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}

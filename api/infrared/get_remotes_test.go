package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestGetRemotes(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetRemotesResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				InfraredID: deviceID,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRemotes(tt.args.InfraredID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRemotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}

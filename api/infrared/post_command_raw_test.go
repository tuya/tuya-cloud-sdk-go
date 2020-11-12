package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestPostCommandRaw(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
		RemoteID   int
		RawKey     int
	}
	tests := []struct {
		name    string
		args    args
		want    *PostCommandRawResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, RemoteID: 5, RawKey: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostCommandRaw(tt.args.InfraredID, tt.args.RemoteID, tt.args.RawKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostCommandRaw() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}

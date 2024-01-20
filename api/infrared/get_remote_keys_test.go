package infrared

import (
	"testing"

	"github.com/tuya/tuya-cloud-sdk-go/api/common"
)

func TestRemoteKeys(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
		RemoteID   string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetRemoteKeysResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, RemoteID: "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetRemoteKeys(tt.args.InfraredID, tt.args.RemoteID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRemoteKeys error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
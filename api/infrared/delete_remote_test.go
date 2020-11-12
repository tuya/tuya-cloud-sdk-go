package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestDeleteRemote(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
		RemoteID   string
	}
	tests := []struct {
		name    string
		args    args
		want    *DeleteRemoteResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{
				InfraredID: deviceID,
				RemoteID:   "6cedcyyyyyyy",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DeleteRemote(tt.args.InfraredID, tt.args.RemoteID)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteRemote() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}

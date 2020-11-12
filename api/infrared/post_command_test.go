package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestPostCommand(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
		RemoteID   int
		Key        string
	}
	tests := []struct {
		name    string
		args    args
		want    *PostCommandResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, RemoteID: 5, Key: "power"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostCommand(tt.args.InfraredID, tt.args.RemoteID, tt.args.Key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}

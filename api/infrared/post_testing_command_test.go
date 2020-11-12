package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestPostTestingCommand(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID  string
		RemoteIndex int
		CategoryID  int
		Key         string
	}
	tests := []struct {
		name    string
		args    args
		want    *PostTestingCommandResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, RemoteIndex: 5129, CategoryID: 5, Key: "power"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostTestingCommand(tt.args.InfraredID, tt.args.RemoteIndex, tt.args.CategoryID, tt.args.Key)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostTestingCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}

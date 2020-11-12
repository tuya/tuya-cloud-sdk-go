package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestPostTestingACCommand(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID  string
		RemoteIndex int
		CategoryID  int
		Code        string
		Value       int
	}
	tests := []struct {
		name    string
		args    args
		want    *PostTestingACCommandResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, RemoteIndex: 7273, CategoryID: 5, Code: "power", Value: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostTestingACCommand(tt.args.InfraredID, tt.args.RemoteIndex, tt.args.CategoryID, tt.args.Code, tt.args.Value)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostTestingACCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}

package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestPostACCommand(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
		RemoteID   string
		Code       string
		Value      int
	}
	tests := []struct {
		name    string
		args    args
		want    *PostACCommandResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, RemoteID: "asdasddasd", Code: "power", Value: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostACCommand(tt.args.InfraredID, tt.args.RemoteID, tt.args.Code, tt.args.Value)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostACCommand() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}

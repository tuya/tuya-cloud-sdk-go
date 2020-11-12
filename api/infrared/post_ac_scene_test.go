package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestPostACScene(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
		RemoteID   string
		Power      int
		Mode       int
		Temp       int
		Wind       int
	}
	tests := []struct {
		name    string
		args    args
		want    *PostACSceneResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, RemoteID: "asdasddasd", Power: 1, Mode: 0, Temp: 19, Wind: 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PostACScene(tt.args.InfraredID, tt.args.RemoteID, tt.args.Power, tt.args.Mode, tt.args.Temp, tt.args.Wind)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostACScene() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}

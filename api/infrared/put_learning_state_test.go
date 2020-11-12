package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestPutLearningState(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
		State      bool
	}
	tests := []struct {
		name    string
		args    args
		want    *PutLearningStateResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID, State: true},
		},
		{
			name: "2",
			args: args{InfraredID: deviceID, State: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PutLearningState(tt.args.InfraredID, tt.args.State)
			if (err != nil) != tt.wantErr {
				t.Errorf("PutLearningState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}

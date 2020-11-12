package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestGetLearningCodes(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
		Time       int64
	}
	tests := []struct {
		name    string
		args    args
		want    *GetLearningCodesResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pls, err := PutLearningState(tt.args.InfraredID, true)
			if err != nil {
				t.Errorf("PutLearningState() error = %v", err)
				return
			}

			got, err := GetLearningCodes(tt.args.InfraredID, pls.T)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLearningCodes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
			_, err = PutLearningState(tt.args.InfraredID, false)
			if err != nil {
				t.Errorf("PutLearningState() error = %v", err)
				return
			}
		})
	}
}

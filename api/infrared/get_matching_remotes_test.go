package infrared

import (
	"testing"

	"github.com/TuyaInc/tuya_cloud_sdk_go/api/common"
)

func TestGetMatchingRemotes(t *testing.T) {
	deviceID := common.Ed.TestDataDeviceID
	type args struct {
		InfraredID string
		Token      string
	}
	tests := []struct {
		name    string
		args    args
		want    *GetMatchingRemotesResponse
		wantErr bool
	}{
		{
			name: "1",
			args: args{InfraredID: deviceID},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenResponse, err := PostMatchingRemotesToken(tt.args.InfraredID, 5, "", "xxxxx")
			if err != nil {
				t.Errorf("PutLearningState() error = %v", err)
				return
			}

			got, err := GetMatchingRemotes(tt.args.InfraredID, tokenResponse.Result.Token)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMatchingRemotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && !got.Success {
				t.Errorf("got no success, msg: %s\n", got.Msg)
			}
			t.Logf("%+v", got)
		})
	}
}
